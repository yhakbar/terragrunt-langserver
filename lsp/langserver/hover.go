package langserver

import (
	"context"
	"github.com/mightyguava/hcl-langserver/lsp/document"
	"github.com/mightyguava/hcl-langserver/lsp/protocol"
	"reflect"
)

// HoverHandler handles hover requests from the LSP client.
type HoverHandler struct {
	w *document.Workspace
}

// NewHoverHandler allocates and initializes a new HoverHandler for the document.Workspace.
func NewHoverHandler(w *document.Workspace) *HoverHandler {
	return &HoverHandler{w}
}

// Handle hover requests from the LSP client.
func (h *HoverHandler) Handle(ctx context.Context, params *protocol.HoverParams) (*protocol.Hover, error) {
	doc, err := h.w.LoadDocument(string(params.TextDocument.URI), true)
	if err != nil {
		return nil, err
	}
	hoverPos := document.ToHclPos(params.Position)
	closest := doc.AST.FindNodeAt(hoverPos)
	if closest == nil {
		return nil, nil
	}
	return &protocol.Hover{
		Contents: protocol.MarkupContent{
			Kind:  protocol.PlainText,
			Value: reflect.TypeOf(closest.Node).String(),
		},
		Range: document.FromHCLRange(closest.Range()),
	}, nil
}
