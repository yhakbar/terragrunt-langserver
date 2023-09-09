package langserver

import (
	"context"
	"github.com/creachadair/jrpc2"
	"github.com/hashicorp/hcl/v2"
	"github.com/mightyguava/terraform-langserver/lang/terragrunt"
	"github.com/mightyguava/terraform-langserver/lsp/protocol"
	"net/url"
	"os"
	"reflect"
)

type HoverHandler struct{}

func (h *HoverHandler) Handle(ctx context.Context, params *protocol.HoverParams) (*protocol.Hover, error) {
	uri, err := url.Parse(string(params.TextDocument.URI))
	if err != nil {
		return nil, jrpc2.Errorf(jrpc2.InvalidParams, "invalid uri %s: %s", params.TextDocument.URI, err.Error())
	}
	if uri.Scheme != "file" {
		return nil, jrpc2.Errorf(jrpc2.InvalidParams, "unsupported scheme %s", uri.Scheme)
	}
	contents, err := os.ReadFile(uri.Path)
	if err != nil {
		return nil, jrpc2.Errorf(jrpc2.InvalidParams, "error opening file %s: %s", params.TextDocument.URI, err.Error())
	}
	ast, err := terragrunt.ParseHCL(uri.Path, contents)
	if err != nil {
		return nil, jrpc2.Errorf(jrpc2.InvalidParams, "error parsing file %s: %s", params.TextDocument.URI, err.Error())
	}
	index := terragrunt.IndexAST(ast)
	hoverPos := ToHclPos(params.Position)
	nodes, ok := index[hoverPos.Line]
	if !ok {
		return nil, nil
	}
	var closest *terragrunt.IndexedNode
	for _, node := range nodes {
		if node.SrcRange.Start.Column < hoverPos.Column {
			closest = node
		}
	}
	return &protocol.Hover{
		Contents: protocol.MarkupContent{
			Kind:  protocol.PlainText,
			Value: reflect.TypeOf(closest.Node).String(),
		},
		Range: FromHCLRange(closest.SrcRange),
	}, nil
}

func FromHCLRange(s *hcl.Range) protocol.Range {
	return protocol.Range{FromHCLPos(s.Start), FromHCLPos(s.End)}
}

func FromHCLPos(s hcl.Pos) protocol.Position {
	return protocol.Position{uint32(max(s.Line-1, 0)), uint32(max(s.Column-1, 0))}
}

func ToHCLRange(s protocol.Range) *hcl.Range {
	return &hcl.Range{"", ToHclPos(s.Start), ToHclPos(s.End)}
}

func ToHclPos(s protocol.Position) hcl.Pos {
	return hcl.Pos{int(s.Line + 1), int(s.Character + 1), 0}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
