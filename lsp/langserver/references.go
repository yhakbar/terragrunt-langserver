package langserver

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/mightyguava/hcl-langserver/lang/terragrunt"
	"github.com/mightyguava/hcl-langserver/lsp/document"
	"github.com/mightyguava/hcl-langserver/lsp/protocol"
	"log/slog"
	"reflect"
)

// Referencer handles reference requests from the LSP client.
type Referencer struct {
	w *document.Workspace
}

// NewReferencer allocates and initializes a new Referencer for the document.Workspace.
func NewReferencer(w *document.Workspace) *Referencer {
	return &Referencer{w}
}

// GoToDefinition handles go to definition requests from the LSP client. For now, it only supports going to definition
// of locals.
func (r *Referencer) GoToDefinition(params protocol.TextDocumentPositionParams) (protocol.Declaration, error) {
	doc, err := r.w.LoadDocument(string(params.TextDocument.URI), true)
	if err != nil {
		return nil, err
	}
	pos := document.ToHclPos(params.Position)
	node := doc.AST.FindNodeAt(pos)
	if node == nil {
		return nil, nil
	}
	slog.Info("Found node", slog.Any("node", node))
	scopeTraversalNode, ok := node.Node.(*hclsyntax.ScopeTraversalExpr)
	if !ok {
		slog.Info("Not scope traversal", "node", node.GoString())
		return nil, nil
	}
	slog.Info("Locals", "v", doc.AST.Locals)
	traversal := scopeTraversalNode.Traversal
	var scope terragrunt.Scope
	for _, t := range traversal {
		slog.Info("Traversing", "type", reflect.TypeOf(t).String())
		switch tv := t.(type) {
		case hcl.TraverseRoot:
			rootName := t.(hcl.TraverseRoot).Name
			if rootName == "local" {
				scope = doc.AST.Locals
			} else {
				slog.Info("scope name not supported", slog.String("scope", rootName))
				return nil, nil
			}
		case hcl.TraverseAttr:
			declaration, ok := scope[tv.Name]
			if ok {
				return []protocol.Location{
					{
						URI:   params.TextDocument.URI,
						Range: document.FromHCLRange(declaration.Range()),
					},
				}, nil
			}
		}
	}
	return nil, nil
}
