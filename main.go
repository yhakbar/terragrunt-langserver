package main

import (
	"context"
	"github.com/creachadair/jrpc2/server"
	"github.com/mightyguava/terraform-langserver/lsp/langserver"
	"github.com/mightyguava/terraform-langserver/lsp/protocol"
	"log"
	"log/slog"
	"net"

	"github.com/creachadair/jrpc2"
	"github.com/creachadair/jrpc2/channel"
)

func main() {
	serviceAddr := "127.0.0.1:8080"
	// -------------------------------------------------------------------------
	// Start the server listening on the local network.
	listener, err := net.Listen(jrpc2.Network(serviceAddr))
	if err != nil {
		log.Fatalf("Listen %q: %v", serviceAddr, err)
	}
	defer listener.Close()

	// -------------------------------------------------------------------------
	// Set up a service with some trivial methods (handlers defined below).
	svc := &langserver.Server{}
	assigner := protocol.SingleAssigner(protocol.JRPCHandler(svc))
	ctx := context.Background()

	slog.Info("Starting server", slog.String("bind", serviceAddr))
	if err := server.Loop(ctx, server.NetAccepter(listener, channel.Line), server.Static(assigner), nil); err != nil {
		log.Fatalf("%+v", err)
	}
}
