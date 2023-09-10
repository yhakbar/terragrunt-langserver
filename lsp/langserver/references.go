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

type Referencer struct {
	w *document.Workspace
}

func NewReferencer(w *document.Workspace) *Referencer {
	return &Referencer{w}
}

func (r *Referencer) GoToDeclaration(params *protocol.DeclarationParams) (protocol.Declaration, error) {
	doc, err := r.w.LoadFile(string(params.TextDocument.URI))
	if err != nil {
		return nil, err
	}
	pos := document.ToHclPos(params.Position)
	node := doc.FindNodeAt(pos)
	if node == nil {
		return nil, nil
	}
	slog.Info("Found node", slog.Any("node", node))
	scopeTraversalNode, ok := node.Node.(*hclsyntax.ScopeTraversalExpr)
	if !ok {
		slog.Info("Not scope traversal, was %s", node.GoString())
		return nil, nil
	}
	slog.Info("Locals", "v", doc.Locals)
	traversal := scopeTraversalNode.Traversal
	var root map[string]*terragrunt.IndexedNode
	for _, t := range traversal {
		slog.Info("Traversing", "type", reflect.TypeOf(t).String())
		switch tv := t.(type) {
		case hcl.TraverseRoot:
			rootName := t.(hcl.TraverseRoot).Name
			if rootName == "local" {
				root = doc.Locals
			} else {
				slog.Info("root name not supported", slog.String("root", rootName))
				return nil, nil
			}
		case hcl.TraverseAttr:
			declaration, ok := root[tv.Name]
			if ok {
				return protocol.Declaration{
					{
						URI:   params.TextDocument.URI,
						Range: document.FromHCLRange(declaration.SrcRange),
					},
				}, nil
			}
		}
	}
	return nil, nil
}
