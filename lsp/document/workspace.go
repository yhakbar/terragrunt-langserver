package document

import (
	"github.com/creachadair/jrpc2"
	"github.com/hashicorp/hcl/v2"
	"github.com/mightyguava/terragrunt-langserver/lang/terragrunt"
	"github.com/mightyguava/terragrunt-langserver/lsp/protocol"
	"github.com/pkg/errors"
	"log/slog"
	"net/url"
	"os"
	"sync"
)

type Document struct {
	// AST contains the AST and index for the document
	AST *terragrunt.IndexedAST
	// TerragruntEval contains the evaluated Terragrunt config
	TerragruntEval *terragrunt.EvaluatedData
	// Diagnostics contains any errors from evaluating the config
	Diagnostics hcl.Diagnostics
}

type Workspace struct {
	openDocs map[string]*Document
	mu       sync.Mutex
}

func NewWorkspace() *Workspace {
	return &Workspace{openDocs: make(map[string]*Document)}
}

func (w *Workspace) LoadDocument(fileUri string, isCachedOk bool) (*Document, error) {
	if isCachedOk {
		w.mu.Lock()
		f, ok := w.openDocs[fileUri]
		w.mu.Unlock()
		if ok {
			return f, nil
		}
	}

	uri, err := parseUri(fileUri)
	if err != nil {
		return nil, err
	}
	contents, err := os.ReadFile(uri.Path)
	if err != nil {
		return nil, jrpc2.Errorf(jrpc2.InvalidParams, "error opening file %s: %s", uri, err.Error())
	}
	return w.LoadDocumentBytes(fileUri, contents)
}

func (w *Workspace) LoadDocumentBytes(fileUri string, contents []byte) (*Document, error) {
	uri, err := parseUri(fileUri)
	if err != nil {
		return nil, err
	}
	file, err := terragrunt.ParseHCLFile(uri.Path, contents)
	var diagnostics hcl.Diagnostics
	if err != nil && !errors.As(err, &diagnostics) {
		return nil, jrpc2.Errorf(jrpc2.InvalidParams, "error parsing file %s: %s", fileUri, err.Error())
	}
	var eval *terragrunt.EvaluatedData
	// If parsing the AST failed, skip evaluation.
	if diagnostics == nil {
		eval, err = terragrunt.Evaluate(uri.Path, file.HCLFile, contents)
		if err != nil {
			if errors.As(err, &diagnostics) {
				// continue
			} else {
				return nil, err
			}
		}
	}
	doc := &Document{
		AST:            file,
		TerragruntEval: eval,
		Diagnostics:    diagnostics,
	}

	w.mu.Lock()
	w.openDocs[fileUri] = doc
	w.mu.Unlock()

	slog.Info("Loaded", slog.String("file", fileUri))

	return doc, nil
}

func (w *Workspace) UnloadFile(fileUri string) {
	w.mu.Lock()
	delete(w.openDocs, fileUri)
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

func FromHCLRange(s hcl.Range) protocol.Range {
	return protocol.Range{FromHCLPos(s.Start), FromHCLPos(s.End)}
}

func FromHCLPos(s hcl.Pos) protocol.Position {
	return protocol.Position{uint32(max(s.Line-1, 0)), uint32(max(s.Column-1, 0))}
}

func ToHCLRange(s protocol.Range) hcl.Range {
	return hcl.Range{"", ToHclPos(s.Start), ToHclPos(s.End)}
}

func ToHclPos(s protocol.Position) hcl.Pos {
	return hcl.Pos{int(s.Line + 1), int(s.Character + 1), 0}
}

func FromHCLDiagnostics(diags hcl.Diagnostics) []protocol.Diagnostic {
	pdiags := make([]protocol.Diagnostic, len(diags))
	for i, hdiag := range diags {
		msg := hdiag.Summary
		if hdiag.Detail != "" {
			msg = hdiag.Detail
		}
		pdiags[i] = protocol.Diagnostic{
			Range:    FromHCLRange(*hdiag.Subject),
			Severity: protocol.DiagnosticSeverity(hdiag.Severity),
			Source:   "terragrunt",
			Message:  msg,
		}
	}
	return pdiags
}
