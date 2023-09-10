package main

import (
	"context"
	"flag"
	"github.com/creachadair/jrpc2"
	"github.com/creachadair/jrpc2/channel"
	"github.com/creachadair/jrpc2/server"
	"github.com/mightyguava/terraform-langserver/lsp/document"
	"github.com/mightyguava/terraform-langserver/lsp/langserver"
	"github.com/mightyguava/terraform-langserver/lsp/protocol"
	"log"
	"log/slog"
	"net"
	"os"
)

type RequestLogger struct{}

var _ jrpc2.RPCLogger = &RequestLogger{}

func (r RequestLogger) LogRequest(ctx context.Context, req *jrpc2.Request) {
	slog.Debug("Request",
		slog.String("id", req.ID()),
		slog.String("method", req.Method()),
		slog.String("params", req.ParamString()),
	)
}

func (r RequestLogger) LogResponse(ctx context.Context, rsp *jrpc2.Response) {
	if rsp.Error() != nil {
		slog.Debug("Response",
			slog.String("id", rsp.ID()),
			slog.String("error", rsp.Error().Error()),
		)
	} else {
		slog.Debug("Response",
			slog.String("id", rsp.ID()),
			slog.String("result", rsp.ResultString()),
		)
	}
}

func main() {
	port := flag.String("socket", "", "port to listen on")
	logLevel := flag.String("log-level", "debug", "log level")
	flag.Parse()
	var level slog.Level
	if err := level.UnmarshalText([]byte(*logLevel)); err != nil {
		log.Fatalf("invalid level %s", *logLevel)
	}
	log.SetFlags(log.Lshortfile)
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: level,
	})))
	// -------------------------------------------------------------------------
	// Set up a service with some trivial methods (handlers defined below).
	workspace := document.NewWorkspace()
	svc := &langserver.Server{
		HoverHandler: langserver.NewHoverHandler(workspace),
		Workspace:    workspace,
		Referencer:   langserver.NewReferencer(workspace),
	}
	assigner := protocol.SingleAssigner(protocol.JRPCHandler(svc))

	serverOpts := &jrpc2.ServerOptions{
		AllowPush: true,
		Logger:    func(text string) { slog.Debug(text) },
		RPCLog:    &RequestLogger{},
	}
	if *port == "" {
		slog.Info("Starting server as stdio")
		srv := jrpc2.NewServer(assigner, serverOpts)
		srv.Start(channel.LSP(os.Stdin, os.Stdout))
		if err := srv.Wait(); err != nil {
			log.Fatalf("%+v", err)
		}
	} else {
		serviceAddr := "127.0.0.1:" + *port

		// -------------------------------------------------------------------------
		// Start the server listening on the local network.
		listener, err := net.Listen(jrpc2.Network(serviceAddr))
		if err != nil {
			log.Fatalf("Listen %q: %v", serviceAddr, err)
		}
		defer listener.Close()

		slog.Info("Starting server", slog.String("bind", serviceAddr))
		if err := server.Loop(context.Background(), server.NetAccepter(listener, channel.LSP), server.Static(assigner), &server.LoopOptions{ServerOptions: serverOpts}); err != nil {
			log.Fatalf("%+v", err)
		}
	}
}
