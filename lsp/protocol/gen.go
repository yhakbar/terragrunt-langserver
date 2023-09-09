package protocol

import (
	"context"
	"github.com/creachadair/jrpc2"
)

//go:generate go run ./gen

type replyAdapter struct {
	result interface{}
}

func (r *replyAdapter) reply(ctx context.Context, result interface{}, err error) error {
	r.result = result
	return err
}

type Replier func(ctx context.Context, result interface{}, err error) error

type JRPCAssigner struct {
	Server
}

var _ jrpc2.Assigner = &JRPCAssigner{}

func (a *JRPCAssigner) Assign(ctx context.Context, method string) jrpc2.Handler {
	return JRPCHandler(a.Server)
}

func JRPCHandler(server Server) func(context.Context, *jrpc2.Request) (any, error) {
	return func(ctx context.Context, request *jrpc2.Request) (any, error) {
		replier := &replyAdapter{}
		isMatched, err := serverDispatch(ctx, server, replier.reply, request)
		if err != nil {
			return nil, err
		}
		if !isMatched {
			return nil, &jrpc2.Error{Code: jrpc2.MethodNotFound, Message: jrpc2.MethodNotFound.String()}
		}
		return replier.result, nil
	}
}

func sendParseError(_ context.Context, _ Replier, err error) error {
	return &jrpc2.Error{Code: jrpc2.ParseError, Message: err.Error()}
}
