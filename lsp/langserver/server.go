package langserver

import (
	"context"
	"github.com/creachadair/jrpc2"
	"github.com/mightyguava/terragrunt-langserver/lsp/document"
	"github.com/mightyguava/terragrunt-langserver/lsp/protocol"
)

// Server implements the LSP protocol's server interface.
type Server struct {
	HoverHandler *HoverHandler
	Workspace    *document.Workspace
	Referencer   *Referencer
	// client provides the server a handle to push messages to the LSP client.
	client protocol.Client
}

// This implementation assertion doesn't actually help, as far as I can tell.
// We're confirming that the server implements gopls' protocol.Server interface.
//
// That requires leaving behind a bunch of dead code for the unimplemented methods.
// It makes the code harder to read and understand.
var _ protocol.Server = &Server{}

// Initialize implements the initialize request. It is the first request from the LSP client to the server.
// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#initialize
func (s *Server) Initialize(ctx context.Context, initialize *protocol.ParamInitialize) (*protocol.InitializeResult, error) {
	s.client = protocol.ClientCaller(protocol.NewSender(jrpc2.ServerFromContext(ctx)))

	return &protocol.InitializeResult{
		ServerInfo: &protocol.ServerInfo{Name: "terragrunt-langserver"},
		Capabilities: protocol.ServerCapabilities{
			HoverProvider: &protocol.Or_ServerCapabilities_hoverProvider{
				Value: true,
			},
			DeclarationProvider: &protocol.Or_ServerCapabilities_declarationProvider{Value: true},
			DefinitionProvider:  &protocol.Or_ServerCapabilities_definitionProvider{Value: true},
			TextDocumentSync: &protocol.TextDocumentSyncOptions{
				OpenClose:         true,
				Change:            protocol.None,
				WillSave:          false,
				WillSaveWaitUntil: false,
				Save:              &protocol.SaveOptions{},
			},
		},
	}, nil
}

func (s *Server) Initialized(ctx context.Context, params *protocol.InitializedParams) error {
	return nil
}

func (s *Server) Progress(ctx context.Context, params *protocol.ProgressParams) error {
	//TODO implement me
	panic("implement me")
}

func (s *Server) SetTrace(ctx context.Context, params *protocol.SetTraceParams) error {
	// TODO: actually set the trace level
	//    https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#traceValue
	return nil
}

func (s *Server) IncomingCalls(ctx context.Context, params *protocol.CallHierarchyIncomingCallsParams) ([]protocol.CallHierarchyIncomingCall, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) OutgoingCalls(ctx context.Context, params *protocol.CallHierarchyOutgoingCallsParams) ([]protocol.CallHierarchyOutgoingCall, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) ResolveCodeAction(ctx context.Context, action *protocol.CodeAction) (*protocol.CodeAction, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) ResolveCodeLens(ctx context.Context, lens *protocol.CodeLens) (*protocol.CodeLens, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) ResolveCompletionItem(ctx context.Context, item *protocol.CompletionItem) (*protocol.CompletionItem, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) ResolveDocumentLink(ctx context.Context, link *protocol.DocumentLink) (*protocol.DocumentLink, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) Exit(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (s *Server) Resolve(ctx context.Context, hint *protocol.InlayHint) (*protocol.InlayHint, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) DidChangeNotebookDocument(ctx context.Context, params *protocol.DidChangeNotebookDocumentParams) error {
	//TODO implement me
	panic("implement me")
}

func (s *Server) DidCloseNotebookDocument(ctx context.Context, params *protocol.DidCloseNotebookDocumentParams) error {
	//TODO implement me
	panic("implement me")
}

func (s *Server) DidOpenNotebookDocument(ctx context.Context, params *protocol.DidOpenNotebookDocumentParams) error {
	//TODO implement me
	panic("implement me")
}

func (s *Server) DidSaveNotebookDocument(ctx context.Context, params *protocol.DidSaveNotebookDocumentParams) error {
	//TODO implement me
	panic("implement me")
}

func (s *Server) Shutdown(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (s *Server) CodeAction(ctx context.Context, params *protocol.CodeActionParams) ([]protocol.CodeAction, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) CodeLens(ctx context.Context, params *protocol.CodeLensParams) ([]protocol.CodeLens, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) ColorPresentation(ctx context.Context, params *protocol.ColorPresentationParams) ([]protocol.ColorPresentation, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) Completion(ctx context.Context, params *protocol.CompletionParams) (*protocol.CompletionList, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) Declaration(ctx context.Context, params *protocol.DeclarationParams) (*protocol.Or_textDocument_declaration, error) {
	result, err := s.Referencer.GoToDefinition(params.TextDocumentPositionParams)
	if err != nil {
		return nil, err
	}

	return &protocol.Or_textDocument_declaration{Value: result}, nil
}

func (s *Server) Definition(ctx context.Context, params *protocol.DefinitionParams) ([]protocol.Location, error) {
	result, err := s.Referencer.GoToDefinition(params.TextDocumentPositionParams)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *Server) Diagnostic(ctx context.Context, params *protocol.DocumentDiagnosticParams) (*protocol.DocumentDiagnosticReport, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) DidChange(ctx context.Context, params *protocol.DidChangeTextDocumentParams) error {
	// TODO: implement me
	return nil
}

func (s *Server) DidClose(ctx context.Context, params *protocol.DidCloseTextDocumentParams) error {
	s.Workspace.UnloadFile(string(params.TextDocument.URI))
	return nil
}

func (s *Server) DidOpen(ctx context.Context, params *protocol.DidOpenTextDocumentParams) error {
	doc, err := s.Workspace.LoadDocumentBytes(string(params.TextDocument.URI), []byte(params.TextDocument.Text))
	if err != nil {
		return err
	}

	_ = s.client.PublishDiagnostics(ctx, &protocol.PublishDiagnosticsParams{
		URI:         params.TextDocument.URI,
		Version:     params.TextDocument.Version,
		Diagnostics: document.FromHCLDiagnostics(doc.Diagnostics),
	})

	return err
}

func (s *Server) DidSave(ctx context.Context, params *protocol.DidSaveTextDocumentParams) error {
	doc, err := s.Workspace.LoadDocument(string(params.TextDocument.URI), false)
	if err != nil {
		return err
	}

	_ = s.client.PublishDiagnostics(ctx, &protocol.PublishDiagnosticsParams{
		URI:         params.TextDocument.URI,
		Diagnostics: document.FromHCLDiagnostics(doc.Diagnostics),
	})

	return err
}

func (s *Server) DocumentColor(ctx context.Context, params *protocol.DocumentColorParams) ([]protocol.ColorInformation, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) DocumentHighlight(ctx context.Context, params *protocol.DocumentHighlightParams) ([]protocol.DocumentHighlight, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) DocumentLink(ctx context.Context, params *protocol.DocumentLinkParams) ([]protocol.DocumentLink, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) DocumentSymbol(ctx context.Context, params *protocol.DocumentSymbolParams) ([]interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) FoldingRange(ctx context.Context, params *protocol.FoldingRangeParams) ([]protocol.FoldingRange, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) Formatting(ctx context.Context, params *protocol.DocumentFormattingParams) ([]protocol.TextEdit, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) Hover(ctx context.Context, params *protocol.HoverParams) (*protocol.Hover, error) {
	return s.HoverHandler.Handle(ctx, params)
}

func (s *Server) Implementation(ctx context.Context, params *protocol.ImplementationParams) ([]protocol.Location, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) InlayHint(ctx context.Context, params *protocol.InlayHintParams) ([]protocol.InlayHint, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) InlineCompletion(ctx context.Context, params *protocol.InlineCompletionParams) (*protocol.Or_Result_textDocument_inlineCompletion, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) InlineValue(ctx context.Context, params *protocol.InlineValueParams) ([]protocol.InlineValue, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) LinkedEditingRange(ctx context.Context, params *protocol.LinkedEditingRangeParams) (*protocol.LinkedEditingRanges, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) Moniker(ctx context.Context, params *protocol.MonikerParams) ([]protocol.Moniker, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) OnTypeFormatting(ctx context.Context, params *protocol.DocumentOnTypeFormattingParams) ([]protocol.TextEdit, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) PrepareCallHierarchy(ctx context.Context, params *protocol.CallHierarchyPrepareParams) ([]protocol.CallHierarchyItem, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) PrepareRename(ctx context.Context, params *protocol.PrepareRenameParams) (*protocol.PrepareRenameResult, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) PrepareTypeHierarchy(ctx context.Context, params *protocol.TypeHierarchyPrepareParams) ([]protocol.TypeHierarchyItem, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) RangeFormatting(ctx context.Context, params *protocol.DocumentRangeFormattingParams) ([]protocol.TextEdit, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) RangesFormatting(ctx context.Context, params *protocol.DocumentRangesFormattingParams) ([]protocol.TextEdit, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) References(ctx context.Context, params *protocol.ReferenceParams) ([]protocol.Location, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) Rename(ctx context.Context, params *protocol.RenameParams) (*protocol.WorkspaceEdit, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) SelectionRange(ctx context.Context, params *protocol.SelectionRangeParams) ([]protocol.SelectionRange, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) SemanticTokensFull(ctx context.Context, params *protocol.SemanticTokensParams) (*protocol.SemanticTokens, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) SemanticTokensFullDelta(ctx context.Context, params *protocol.SemanticTokensDeltaParams) (interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) SemanticTokensRange(ctx context.Context, params *protocol.SemanticTokensRangeParams) (*protocol.SemanticTokens, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) SignatureHelp(ctx context.Context, params *protocol.SignatureHelpParams) (*protocol.SignatureHelp, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) TypeDefinition(ctx context.Context, params *protocol.TypeDefinitionParams) ([]protocol.Location, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) WillSave(ctx context.Context, params *protocol.WillSaveTextDocumentParams) error {
	//TODO implement me
	panic("implement me")
}

func (s *Server) WillSaveWaitUntil(ctx context.Context, params *protocol.WillSaveTextDocumentParams) ([]protocol.TextEdit, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) Subtypes(ctx context.Context, params *protocol.TypeHierarchySubtypesParams) ([]protocol.TypeHierarchyItem, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) Supertypes(ctx context.Context, params *protocol.TypeHierarchySupertypesParams) ([]protocol.TypeHierarchyItem, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) WorkDoneProgressCancel(ctx context.Context, params *protocol.WorkDoneProgressCancelParams) error {
	//TODO implement me
	panic("implement me")
}

func (s *Server) DiagnosticWorkspace(ctx context.Context, params *protocol.WorkspaceDiagnosticParams) (*protocol.WorkspaceDiagnosticReport, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) DidChangeConfiguration(ctx context.Context, params *protocol.DidChangeConfigurationParams) error {
	//TODO implement me
	panic("implement me")
}

func (s *Server) DidChangeWatchedFiles(ctx context.Context, params *protocol.DidChangeWatchedFilesParams) error {
	//TODO implement me
	panic("implement me")
}

func (s *Server) DidChangeWorkspaceFolders(ctx context.Context, params *protocol.DidChangeWorkspaceFoldersParams) error {
	//TODO implement me
	panic("implement me")
}

func (s *Server) DidCreateFiles(ctx context.Context, params *protocol.CreateFilesParams) error {
	//TODO implement me
	panic("implement me")
}

func (s *Server) DidDeleteFiles(ctx context.Context, params *protocol.DeleteFilesParams) error {
	//TODO implement me
	panic("implement me")
}

func (s *Server) DidRenameFiles(ctx context.Context, params *protocol.RenameFilesParams) error {
	//TODO implement me
	panic("implement me")
}

func (s *Server) ExecuteCommand(ctx context.Context, params *protocol.ExecuteCommandParams) (interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) Symbol(ctx context.Context, params *protocol.WorkspaceSymbolParams) ([]protocol.SymbolInformation, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) WillCreateFiles(ctx context.Context, params *protocol.CreateFilesParams) (*protocol.WorkspaceEdit, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) WillDeleteFiles(ctx context.Context, params *protocol.DeleteFilesParams) (*protocol.WorkspaceEdit, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) WillRenameFiles(ctx context.Context, params *protocol.RenameFilesParams) (*protocol.WorkspaceEdit, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) ResolveWorkspaceSymbol(ctx context.Context, symbol *protocol.WorkspaceSymbol) (*protocol.WorkspaceSymbol, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) NonstandardRequest(ctx context.Context, method string, params interface{}) (interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) TextDocumentContent(ctx context.Context, params *protocol.TextDocumentContentParams) (*string, error) {
	//TODO implement me
	panic("implement me")
}
