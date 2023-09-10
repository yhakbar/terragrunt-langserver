package document

import (
	"github.com/creachadair/jrpc2"
	"github.com/hashicorp/hcl/v2"
	"github.com/mightyguava/terraform-langserver/lang/terragrunt"
	"github.com/mightyguava/terraform-langserver/lsp/protocol"
	"log/slog"
	"net/url"
	"os"
	"sync"
)

type Workspace struct {
	openFiles map[string]*terragrunt.Document
	mu        sync.Mutex
}

func NewWorkspace() *Workspace {
	return &Workspace{openFiles: make(map[string]*terragrunt.Document)}
}

func (w *Workspace) LoadFile(fileUri string) (*terragrunt.Document, error) {
	w.mu.Lock()
	f, ok := w.openFiles[fileUri]
	w.mu.Unlock()
	if ok {
		return f, nil
	}

	uri, err := parseUri(fileUri)
	if err != nil {
		return nil, err
	}
	contents, err := os.ReadFile(uri.Path)
	if err != nil {
		return nil, jrpc2.Errorf(jrpc2.InvalidParams, "error opening file %s: %s", uri, err.Error())
	}
	return w.LoadFileContents(fileUri, contents)
}

func (w *Workspace) LoadFileContents(fileUri string, contents []byte) (*terragrunt.Document, error) {
	ast, err := terragrunt.ParseHCL(fileUri, contents)
	if err != nil {
		return nil, jrpc2.Errorf(jrpc2.InvalidParams, "error parsing file %s: %s", fileUri, err.Error())
	}
	doc := terragrunt.IndexAST(ast)

	w.mu.Lock()
	w.openFiles[fileUri] = doc
	w.mu.Unlock()

	slog.Info("Loaded", slog.String("file", fileUri))

	return doc, nil
}

func (w *Workspace) UnloadFile(fileUri string) {
	w.mu.Lock()
	delete(w.openFiles, fileUri)
	w.mu.Unlock()

	slog.Info("Unloaded", slog.String("file", fileUri))
}

func parseUri(fileUri string) (*url.URL, error) {
	uri, err := url.Parse(fileUri)
	if err != nil {
		return nil, jrpc2.Errorf(jrpc2.InvalidParams, "invalid uri %s: %s", uri, err.Error())
	}
	if uri.Scheme != "file" {
		return nil, jrpc2.Errorf(jrpc2.InvalidParams, "unsupported scheme %s", uri)
	}
	return uri, nil
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
