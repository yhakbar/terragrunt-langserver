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

// SingleAssigner is a jrpc2.Assigner that always assigns to a single static jrpc2.Handler.
type SingleAssigner jrpc2.Handler

var _ jrpc2.Assigner = SingleAssigner(jrpc2.Handler(nil))

func (s SingleAssigner) Assign(ctx context.Context, method string) jrpc2.Handler {
	return s
}

// JRPCHandler adapts the generated protocol code to a jrpc2.Handler
func JRPCHandler(server Server) jrpc2.Handler {
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
