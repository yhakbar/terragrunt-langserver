package langserver

import (
	"context"
	"github.com/creachadair/jrpc2"
	"log/slog"
)

// RequestLogger logs request and responses to the language server.
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
