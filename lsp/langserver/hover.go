package langserver

import (
	"context"
	"github.com/mightyguava/terraform-langserver/lsp/document"
	"github.com/mightyguava/terraform-langserver/lsp/protocol"
	"reflect"
)

type HoverHandler struct {
	w *document.Workspace
}

func NewHoverHandler(w *document.Workspace) *HoverHandler {
	return &HoverHandler{w}
}

func (h *HoverHandler) Handle(ctx context.Context, params *protocol.HoverParams) (*protocol.Hover, error) {
	doc, err := h.w.LoadFile(string(params.TextDocument.URI))
	if err != nil {
		return nil, err
	}
	hoverPos := document.ToHclPos(params.Position)
	closest := doc.FindNodeAt(hoverPos)
	if closest == nil {
		return nil, nil
	}
	return &protocol.Hover{
		Contents: protocol.MarkupContent{
			Kind:  protocol.PlainText,
			Value: reflect.TypeOf(closest.Node).String(),
		},
		Range: document.FromHCLRange(closest.SrcRange),
	}, nil
}
