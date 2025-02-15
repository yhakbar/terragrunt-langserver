package main

import (
	"context"
	"flag"
	"github.com/creachadair/jrpc2"
	"github.com/creachadair/jrpc2/channel"
	"github.com/creachadair/jrpc2/server"
	"github.com/mightyguava/terragrunt-langserver/lsp/document"
	"github.com/mightyguava/terragrunt-langserver/lsp/langserver"
	"github.com/mightyguava/terragrunt-langserver/lsp/protocol"
	"log"
	"log/slog"
	"net"
	"os"
)

func main() {
	port := flag.String("socket", "", "port to listen on")
	stdio := flag.Bool("stdio", false, "use stdio transport")
	logLevel := flag.String("log-level", "debug", "log level")
	logRequests := flag.Bool("log-requests", false, "set to log request payloads")
	debug := flag.Bool("debug", false, "set to enable debug hover hints")

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
		HoverHandler: langserver.NewHoverHandler(workspace, *debug),
		Workspace:    workspace,
		Referencer:   langserver.NewReferencer(workspace),
	}
	assigner := protocol.SingleAssigner(protocol.JRPCHandler(svc))

	var requestLogger jrpc2.RPCLogger
	if *logRequests {
		requestLogger = &langserver.RequestLogger{}
	}
	serverOpts := &jrpc2.ServerOptions{
		AllowPush: true,
		Logger:    func(text string) { slog.Log(nil, slog.Level(-5), text) },
		RPCLog:    requestLogger,
	}
	if *stdio || *port == "" {
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
