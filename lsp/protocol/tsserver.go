// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated for LSP. DO NOT EDIT.

package protocol

// Code generated from protocol/metaModel.json at ref release/protocol/3.17.4-next.2 (hash 184c8a7f010d335582f24337fe182baa6f2fccdd).
// https://github.com/microsoft/vscode-languageserver-node/blob/release/protocol/3.17.4-next.2/protocol/metaModel.json
// LSP metaData.version = 3.17.0.

import (
	"context"
	
	"github.com/creachadair/jrpc2"
)

type Server interface {
	Progress(context.Context, *ProgressParams) error                                                             // $/progress
	SetTrace(context.Context, *SetTraceParams) error                                                             // $/setTrace
	IncomingCalls(context.Context, *CallHierarchyIncomingCallsParams) ([]CallHierarchyIncomingCall, error)       // callHierarchy/incomingCalls
	OutgoingCalls(context.Context, *CallHierarchyOutgoingCallsParams) ([]CallHierarchyOutgoingCall, error)       // callHierarchy/outgoingCalls
	ResolveCodeAction(context.Context, *CodeAction) (*CodeAction, error)                                         // codeAction/resolve
	ResolveCodeLens(context.Context, *CodeLens) (*CodeLens, error)                                               // codeLens/resolve
	ResolveCompletionItem(context.Context, *CompletionItem) (*CompletionItem, error)                             // completionItem/resolve
	ResolveDocumentLink(context.Context, *DocumentLink) (*DocumentLink, error)                                   // documentLink/resolve
	Exit(context.Context) error                                                                                  // exit
	Initialize(context.Context, *ParamInitialize) (*InitializeResult, error)                                     // initialize
	Initialized(context.Context, *InitializedParams) error                                                       // initialized
	Resolve(context.Context, *InlayHint) (*InlayHint, error)                                                     // inlayHint/resolve
	DidChangeNotebookDocument(context.Context, *DidChangeNotebookDocumentParams) error                           // notebookDocument/didChange
	DidCloseNotebookDocument(context.Context, *DidCloseNotebookDocumentParams) error                             // notebookDocument/didClose
	DidOpenNotebookDocument(context.Context, *DidOpenNotebookDocumentParams) error                               // notebookDocument/didOpen
	DidSaveNotebookDocument(context.Context, *DidSaveNotebookDocumentParams) error                               // notebookDocument/didSave
	Shutdown(context.Context) error                                                                              // shutdown
	CodeAction(context.Context, *CodeActionParams) ([]CodeAction, error)                                         // textDocument/codeAction
	CodeLens(context.Context, *CodeLensParams) ([]CodeLens, error)                                               // textDocument/codeLens
	ColorPresentation(context.Context, *ColorPresentationParams) ([]ColorPresentation, error)                    // textDocument/colorPresentation
	Completion(context.Context, *CompletionParams) (*CompletionList, error)                                      // textDocument/completion
	Declaration(context.Context, *DeclarationParams) (*Or_textDocument_declaration, error)                       // textDocument/declaration
	Definition(context.Context, *DefinitionParams) ([]Location, error)                                           // textDocument/definition
	Diagnostic(context.Context, *string) (*string, error)                                                        // textDocument/diagnostic
	DidChange(context.Context, *DidChangeTextDocumentParams) error                                               // textDocument/didChange
	DidClose(context.Context, *DidCloseTextDocumentParams) error                                                 // textDocument/didClose
	DidOpen(context.Context, *DidOpenTextDocumentParams) error                                                   // textDocument/didOpen
	DidSave(context.Context, *DidSaveTextDocumentParams) error                                                   // textDocument/didSave
	DocumentColor(context.Context, *DocumentColorParams) ([]ColorInformation, error)                             // textDocument/documentColor
	DocumentHighlight(context.Context, *DocumentHighlightParams) ([]DocumentHighlight, error)                    // textDocument/documentHighlight
	DocumentLink(context.Context, *DocumentLinkParams) ([]DocumentLink, error)                                   // textDocument/documentLink
	DocumentSymbol(context.Context, *DocumentSymbolParams) ([]interface{}, error)                                // textDocument/documentSymbol
	FoldingRange(context.Context, *FoldingRangeParams) ([]FoldingRange, error)                                   // textDocument/foldingRange
	Formatting(context.Context, *DocumentFormattingParams) ([]TextEdit, error)                                   // textDocument/formatting
	Hover(context.Context, *HoverParams) (*Hover, error)                                                         // textDocument/hover
	Implementation(context.Context, *ImplementationParams) ([]Location, error)                                   // textDocument/implementation
	InlayHint(context.Context, *InlayHintParams) ([]InlayHint, error)                                            // textDocument/inlayHint
	InlineCompletion(context.Context, *InlineCompletionParams) (*Or_Result_textDocument_inlineCompletion, error) // textDocument/inlineCompletion
	InlineValue(context.Context, *InlineValueParams) ([]InlineValue, error)                                      // textDocument/inlineValue
	LinkedEditingRange(context.Context, *LinkedEditingRangeParams) (*LinkedEditingRanges, error)                 // textDocument/linkedEditingRange
	Moniker(context.Context, *MonikerParams) ([]Moniker, error)                                                  // textDocument/moniker
	OnTypeFormatting(context.Context, *DocumentOnTypeFormattingParams) ([]TextEdit, error)                       // textDocument/onTypeFormatting
	PrepareCallHierarchy(context.Context, *CallHierarchyPrepareParams) ([]CallHierarchyItem, error)              // textDocument/prepareCallHierarchy
	PrepareRename(context.Context, *PrepareRenameParams) (*PrepareRename2Gn, error)                              // textDocument/prepareRename
	PrepareTypeHierarchy(context.Context, *TypeHierarchyPrepareParams) ([]TypeHierarchyItem, error)              // textDocument/prepareTypeHierarchy
	RangeFormatting(context.Context, *DocumentRangeFormattingParams) ([]TextEdit, error)                         // textDocument/rangeFormatting
	RangesFormatting(context.Context, *DocumentRangesFormattingParams) ([]TextEdit, error)                       // textDocument/rangesFormatting
	References(context.Context, *ReferenceParams) ([]Location, error)                                            // textDocument/references
	Rename(context.Context, *RenameParams) (*WorkspaceEdit, error)                                               // textDocument/rename
	SelectionRange(context.Context, *SelectionRangeParams) ([]SelectionRange, error)                             // textDocument/selectionRange
	SemanticTokensFull(context.Context, *SemanticTokensParams) (*SemanticTokens, error)                          // textDocument/semanticTokens/full
	SemanticTokensFullDelta(context.Context, *SemanticTokensDeltaParams) (interface{}, error)                    // textDocument/semanticTokens/full/delta
	SemanticTokensRange(context.Context, *SemanticTokensRangeParams) (*SemanticTokens, error)                    // textDocument/semanticTokens/range
	SignatureHelp(context.Context, *SignatureHelpParams) (*SignatureHelp, error)                                 // textDocument/signatureHelp
	TypeDefinition(context.Context, *TypeDefinitionParams) ([]Location, error)                                   // textDocument/typeDefinition
	WillSave(context.Context, *WillSaveTextDocumentParams) error                                                 // textDocument/willSave
	WillSaveWaitUntil(context.Context, *WillSaveTextDocumentParams) ([]TextEdit, error)                          // textDocument/willSaveWaitUntil
	Subtypes(context.Context, *TypeHierarchySubtypesParams) ([]TypeHierarchyItem, error)                         // typeHierarchy/subtypes
	Supertypes(context.Context, *TypeHierarchySupertypesParams) ([]TypeHierarchyItem, error)                     // typeHierarchy/supertypes
	WorkDoneProgressCancel(context.Context, *WorkDoneProgressCancelParams) error                                 // window/workDoneProgress/cancel
	DiagnosticWorkspace(context.Context, *WorkspaceDiagnosticParams) (*WorkspaceDiagnosticReport, error)         // workspace/diagnostic
	DidChangeConfiguration(context.Context, *DidChangeConfigurationParams) error                                 // workspace/didChangeConfiguration
	DidChangeWatchedFiles(context.Context, *DidChangeWatchedFilesParams) error                                   // workspace/didChangeWatchedFiles
	DidChangeWorkspaceFolders(context.Context, *DidChangeWorkspaceFoldersParams) error                           // workspace/didChangeWorkspaceFolders
	DidCreateFiles(context.Context, *CreateFilesParams) error                                                    // workspace/didCreateFiles
	DidDeleteFiles(context.Context, *DeleteFilesParams) error                                                    // workspace/didDeleteFiles
	DidRenameFiles(context.Context, *RenameFilesParams) error                                                    // workspace/didRenameFiles
	ExecuteCommand(context.Context, *ExecuteCommandParams) (interface{}, error)                                  // workspace/executeCommand
	Symbol(context.Context, *WorkspaceSymbolParams) ([]SymbolInformation, error)                                 // workspace/symbol
	WillCreateFiles(context.Context, *CreateFilesParams) (*WorkspaceEdit, error)                                 // workspace/willCreateFiles
	WillDeleteFiles(context.Context, *DeleteFilesParams) (*WorkspaceEdit, error)                                 // workspace/willDeleteFiles
	WillRenameFiles(context.Context, *RenameFilesParams) (*WorkspaceEdit, error)                                 // workspace/willRenameFiles
	ResolveWorkspaceSymbol(context.Context, *WorkspaceSymbol) (*WorkspaceSymbol, error)                          // workspaceSymbol/resolve
	NonstandardRequest(ctx context.Context, method string, params interface{}) (interface{}, error)
}

func serverDispatch(ctx context.Context, server Server, reply Replier, r *jrpc2.Request) (bool, error) {
	switch r.Method() {
	case "$/progress":
		var params ProgressParams
		if err := r.UnmarshalParams(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}
		err := server.Progress(ctx, &params)
		return true, reply(ctx, nil, err)
	case "$/setTrace":
		var params SetTraceParams
		if err := r.UnmarshalParams(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}
		err := server.SetTrace(ctx, &params)
		return true, reply(ctx, nil, err)
	case "callHierarchy/incomingCalls":
		var params CallHierarchyIncomingCallsParams
		if err := r.UnmarshalParams(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}
		resp, err := server.IncomingCalls(ctx, &params)
		if err != nil {
			return true, reply(ctx, nil, err)
		}
		return true, reply(ctx, resp, nil)
	case "callHierarchy/outgoingCalls":
		var params CallHierarchyOutgoingCallsParams
		if err := r.UnmarshalParams(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}
		resp, err := server.OutgoingCalls(ctx, &params)
		if err != nil {
			return true, reply(ctx, nil, err)
		}
		return true, reply(ctx, resp, nil)
	case "codeAction/resolve":
		var params CodeAction
		if err := r.UnmarshalParams(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}
		resp, err := server.ResolveCodeAction(ctx, &params)
		if err != nil {
			return true, reply(ctx, nil, err)
		}
		return true, reply(ctx, resp, nil)
	case "codeLens/resolve":
		var params CodeLens
		if err := r.UnmarshalParams(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}
		resp, err := server.ResolveCodeLens(ctx, &params)
		if err != nil {
			return true, reply(ctx, nil, err)
		}
		return true, reply(ctx, resp, nil)
	case "completionItem/resolve":
		var params CompletionItem
		if err := r.UnmarshalParams(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}
		resp, err := server.ResolveCompletionItem(ctx, &params)
		if err != nil {
			return true, reply(ctx, nil, err)
		}
		return true, reply(ctx, resp, nil)
	case "documentLink/resolve":
		var params DocumentLink
		if err := r.UnmarshalParams(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}
		resp, err := server.ResolveDocumentLink(ctx, &params)
		if err != nil {
			return true, reply(ctx, nil, err)
		}
		return true, reply(ctx, resp, nil)
	case "exit":
		err := server.Exit(ctx)
		return true, reply(ctx, nil, err)
	case "initialize":
		var params ParamInitialize
		if err := r.UnmarshalParams(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}
		resp, err := server.Initialize(ctx, &params)
		if err != nil {
			return true, reply(ctx, nil, err)
		}
		return true, reply(ctx, resp, nil)
	case "initialized":
		var params InitializedParams
		if err := r.UnmarshalParams(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}
		err := server.Initialized(ctx, &params)
		return true, reply(ctx, nil, err)
	case "inlayHint/resolve":
		var params InlayHint
		if err := r.UnmarshalParams(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}
		resp, err := server.Resolve(ctx, &params)
		if err != nil {
			return true, reply(ctx, nil, err)
		}
		return true, reply(ctx, resp, nil)
	case "notebookDocument/didChange":
		var params DidChangeNotebookDocumentParams
		if err := r.UnmarshalParams(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}
		err := server.DidChangeNotebookDocument(ctx, &params)
		return true, reply(ctx, nil, err)
	case "notebookDocument/didClose":
		var params DidCloseNotebookDocumentParams
		if err := r.UnmarshalParams(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}
		err := server.DidCloseNotebookDocument(ctx, &params)
		return true, reply(ctx, nil, err)
	case "notebookDocument/didOpen":
		var params DidOpenNotebookDocumentParams
		if err := r.UnmarshalParams(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}
		err := server.DidOpenNotebookDocument(ctx, &params)
		return true, reply(ctx, nil, err)
	case "notebookDocument/didSave":
		var params DidSaveNotebookDocumentParams
		if err := r.UnmarshalParams(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}
		err := server.DidSaveNotebookDocument(ctx, &params)
		return true, reply(ctx, nil, err)
	case "shutdown":
		err := server.Shutdown(ctx)
		return true, reply(ctx, nil, err)
	case "textDocument/codeAction":
		var params CodeActionParams
		if err := r.UnmarshalParams(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}
		resp, err := server.CodeAction(ctx, &params)
		if err != nil {
			return true, reply(ctx, nil, err)
		}
		return true, reply(ctx, resp, nil)
	case "textDocument/codeLens":
		var params CodeLensParams
		if err := r.UnmarshalParams(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}
		resp, err := server.CodeLens(ctx, &params)
		if err != nil {
			return true, reply(ctx, nil, err)
		}
		return true, reply(ctx, resp, nil)
	case "textDocument/colorPresentation":
		var params ColorPresentationParams
		if err := r.UnmarshalParams(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}
		resp, err := server.ColorPresentation(ctx, &params)
		if err != nil {
			return true, reply(ctx, nil, err)
		}
		return true, reply(ctx, resp, nil)
	case "textDocument/completion":
		var params CompletionParams
		if err := r.UnmarshalParams(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}
		resp, err := server.Completion(ctx, &params)
		if err != nil {
			return true, reply(ctx, nil, err)
		}
		return true, reply(ctx, resp, nil)
	case "textDocument/declaration":
		var params DeclarationParams
		if err := r.UnmarshalParams(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}
		resp, err := server.Declaration(ctx, &params)
		if err != nil {
			return true, reply(ctx, nil, err)
		}
		return true, reply(ctx, resp, nil)
	case "textDocument/definition":
		var params DefinitionParams
		if err := r.UnmarshalParams(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}
		resp, err := server.Definition(ctx, &params)
		if err != nil {
			return true, reply(ctx, nil, err)
		}
		return true, reply(ctx, resp, nil)
	case "textDocument/diagnostic":
		var params string
		if err := r.UnmarshalParams(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}
		resp, err := server.Diagnostic(ctx, &params)
		if err != nil {
			return true, reply(ctx, nil, err)
		}
		return true, reply(ctx, resp, nil)
	case "textDocument/didChange":
		var params DidChangeTextDocumentParams
		if err := r.UnmarshalParams(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}
		err := server.DidChange(ctx, &params)
		return true, reply(ctx, nil, err)
	case "textDocument/didClose":
		var params DidCloseTextDocumentParams
		if err := r.UnmarshalParams(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}
		err := server.DidClose(ctx, &params)
		return true, reply(ctx, nil, err)
	case "textDocument/didOpen":
		var params DidOpenTextDocumentParams
		if err := r.UnmarshalParams(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}
		err := server.DidOpen(ctx, &params)
		return true, reply(ctx, nil, err)
	case "textDocument/didSave":
		var params DidSaveTextDocumentParams
		if err := r.UnmarshalParams(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}
		err := server.DidSave(ctx, &params)
		return true, reply(ctx, nil, err)
	case "textDocument/documentColor":
		var params DocumentColorParams
		if err := r.UnmarshalParams(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}
		resp, err := server.DocumentColor(ctx, &params)
		if err != nil {
			return true, reply(ctx, nil, err)
		}
		return true, reply(ctx, resp, nil)
	case "textDocument/documentHighlight":
		var params DocumentHighlightParams
		if err := r.UnmarshalParams(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}
		resp, err := server.DocumentHighlight(ctx, &params)
		if err != nil {
			return true, reply(ctx, nil, err)
		}
		return true, reply(ctx, resp, nil)
	case "textDocument/documentLink":
		var params DocumentLinkParams
		if err := r.UnmarshalParams(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}
		resp, err := server.DocumentLink(ctx, &params)
		if err != nil {
			return true, reply(ctx, nil, err)
		}
		return true, reply(ctx, resp, nil)
	case "textDocument/documentSymbol":
		var params DocumentSymbolParams
		if err := r.UnmarshalParams(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}
		resp, err := server.DocumentSymbol(ctx, &params)
		if err != nil {
			return true, reply(ctx, nil, err)
		}
		return true, reply(ctx, resp, nil)
	case "textDocument/foldingRange":
		var params FoldingRangeParams
		if err := r.UnmarshalParams(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}
		resp, err := server.FoldingRange(ctx, &params)
		if err != nil {
			return true, reply(ctx, nil, err)
		}
		return true, reply(ctx, resp, nil)
	case "textDocument/formatting":
		var params DocumentFormattingParams
		if err := r.UnmarshalParams(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}
		resp, err := server.Formatting(ctx, &params)
		if err != nil {
			return true, reply(ctx, nil, err)
		}
		return true, reply(ctx, resp, nil)
	case "textDocument/hover":
		var params HoverParams
		if err := r.UnmarshalParams(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}
		resp, err := server.Hover(ctx, &params)
		if err != nil {
			return true, reply(ctx, nil, err)
		}
		return true, reply(ctx, resp, nil)
	case "textDocument/implementation":
		var params ImplementationParams
		if err := r.UnmarshalParams(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}
		resp, err := server.Implementation(ctx, &params)
		if err != nil {
			return true, reply(ctx, nil, err)
		}
		return true, reply(ctx, resp, nil)
	case "textDocument/inlayHint":
		var params InlayHintParams
		if err := r.UnmarshalParams(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}
		resp, err := server.InlayHint(ctx, &params)
		if err != nil {
			return true, reply(ctx, nil, err)
		}
		return true, reply(ctx, resp, nil)
	case "textDocument/inlineCompletion":
		var params InlineCompletionParams
		if err := r.UnmarshalParams(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}
		resp, err := server.InlineCompletion(ctx, &params)
		if err != nil {
			return true, reply(ctx, nil, err)
		}
		return true, reply(ctx, resp, nil)
	case "textDocument/inlineValue":
		var params InlineValueParams
		if err := r.UnmarshalParams(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}
		resp, err := server.InlineValue(ctx, &params)
		if err != nil {
			return true, reply(ctx, nil, err)
		}
		return true, reply(ctx, resp, nil)
	case "textDocument/linkedEditingRange":
		var params LinkedEditingRangeParams
		if err := r.UnmarshalParams(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}
		resp, err := server.LinkedEditingRange(ctx, &params)
		if err != nil {
			return true, reply(ctx, nil, err)
		}
		return true, reply(ctx, resp, nil)
	case "textDocument/moniker":
		var params MonikerParams
		if err := r.UnmarshalParams(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}
		resp, err := server.Moniker(ctx, &params)
		if err != nil {
			return true, reply(ctx, nil, err)
		}
		return true, reply(ctx, resp, nil)
	case "textDocument/onTypeFormatting":
		var params DocumentOnTypeFormattingParams
		if err := r.UnmarshalParams(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}
		resp, err := server.OnTypeFormatting(ctx, &params)
		if err != nil {
			return true, reply(ctx, nil, err)
		}
		return true, reply(ctx, resp, nil)
	case "textDocument/prepareCallHierarchy":
		var params CallHierarchyPrepareParams
		if err := r.UnmarshalParams(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}
		resp, err := server.PrepareCallHierarchy(ctx, &params)
		if err != nil {
			return true, reply(ctx, nil, err)
		}
		return true, reply(ctx, resp, nil)
	case "textDocument/prepareRename":
		var params PrepareRenameParams
		if err := r.UnmarshalParams(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}
		resp, err := server.PrepareRename(ctx, &params)
		if err != nil {
			return true, reply(ctx, nil, err)
		}
		return true, reply(ctx, resp, nil)
	case "textDocument/prepareTypeHierarchy":
		var params TypeHierarchyPrepareParams
		if err := r.UnmarshalParams(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}
		resp, err := server.PrepareTypeHierarchy(ctx, &params)
		if err != nil {
			return true, reply(ctx, nil, err)
		}
		return true, reply(ctx, resp, nil)
	case "textDocument/rangeFormatting":
		var params DocumentRangeFormattingParams
		if err := r.UnmarshalParams(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}
		resp, err := server.RangeFormatting(ctx, &params)
		if err != nil {
			return true, reply(ctx, nil, err)
		}
		return true, reply(ctx, resp, nil)
	case "textDocument/rangesFormatting":
		var params DocumentRangesFormattingParams
		if err := r.UnmarshalParams(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}
		resp, err := server.RangesFormatting(ctx, &params)
		if err != nil {
			return true, reply(ctx, nil, err)
		}
		return true, reply(ctx, resp, nil)
	case "textDocument/references":
		var params ReferenceParams
		if err := r.UnmarshalParams(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}
		resp, err := server.References(ctx, &params)
		if err != nil {
			return true, reply(ctx, nil, err)
		}
		return true, reply(ctx, resp, nil)
	case "textDocument/rename":
		var params RenameParams
		if err := r.UnmarshalParams(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}
		resp, err := server.Rename(ctx, &params)
		if err != nil {
			return true, reply(ctx, nil, err)
		}
		return true, reply(ctx, resp, nil)
	case "textDocument/selectionRange":
		var params SelectionRangeParams
		if err := r.UnmarshalParams(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}
		resp, err := server.SelectionRange(ctx, &params)
		if err != nil {
			return true, reply(ctx, nil, err)
		}
		return true, reply(ctx, resp, nil)
	case "textDocument/semanticTokens/full":
		var params SemanticTokensParams
		if err := r.UnmarshalParams(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}
		resp, err := server.SemanticTokensFull(ctx, &params)
		if err != nil {
			return true, reply(ctx, nil, err)
		}
		return true, reply(ctx, resp, nil)
	case "textDocument/semanticTokens/full/delta":
		var params SemanticTokensDeltaParams
		if err := r.UnmarshalParams(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}
		resp, err := server.SemanticTokensFullDelta(ctx, &params)
		if err != nil {
			return true, reply(ctx, nil, err)
		}
		return true, reply(ctx, resp, nil)
	case "textDocument/semanticTokens/range":
		var params SemanticTokensRangeParams
		if err := r.UnmarshalParams(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}
		resp, err := server.SemanticTokensRange(ctx, &params)
		if err != nil {
			return true, reply(ctx, nil, err)
		}
		return true, reply(ctx, resp, nil)
	case "textDocument/signatureHelp":
		var params SignatureHelpParams
		if err := r.UnmarshalParams(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}
		resp, err := server.SignatureHelp(ctx, &params)
		if err != nil {
			return true, reply(ctx, nil, err)
		}
		return true, reply(ctx, resp, nil)
	case "textDocument/typeDefinition":
		var params TypeDefinitionParams
		if err := r.UnmarshalParams(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}
		resp, err := server.TypeDefinition(ctx, &params)
		if err != nil {
			return true, reply(ctx, nil, err)
		}
		return true, reply(ctx, resp, nil)
	case "textDocument/willSave":
		var params WillSaveTextDocumentParams
		if err := r.UnmarshalParams(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}
		err := server.WillSave(ctx, &params)
		return true, reply(ctx, nil, err)
	case "textDocument/willSaveWaitUntil":
		var params WillSaveTextDocumentParams
		if err := r.UnmarshalParams(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}
		resp, err := server.WillSaveWaitUntil(ctx, &params)
		if err != nil {
			return true, reply(ctx, nil, err)
		}
		return true, reply(ctx, resp, nil)
	case "typeHierarchy/subtypes":
		var params TypeHierarchySubtypesParams
		if err := r.UnmarshalParams(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}
		resp, err := server.Subtypes(ctx, &params)
		if err != nil {
			return true, reply(ctx, nil, err)
		}
		return true, reply(ctx, resp, nil)
	case "typeHierarchy/supertypes":
		var params TypeHierarchySupertypesParams
		if err := r.UnmarshalParams(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}
		resp, err := server.Supertypes(ctx, &params)
		if err != nil {
			return true, reply(ctx, nil, err)
		}
		return true, reply(ctx, resp, nil)
	case "window/workDoneProgress/cancel":
		var params WorkDoneProgressCancelParams
		if err := r.UnmarshalParams(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}
		err := server.WorkDoneProgressCancel(ctx, &params)
		return true, reply(ctx, nil, err)
	case "workspace/diagnostic":
		var params WorkspaceDiagnosticParams
		if err := r.UnmarshalParams(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}
		resp, err := server.DiagnosticWorkspace(ctx, &params)
		if err != nil {
			return true, reply(ctx, nil, err)
		}
		return true, reply(ctx, resp, nil)
	case "workspace/didChangeConfiguration":
		var params DidChangeConfigurationParams
		if err := r.UnmarshalParams(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}
		err := server.DidChangeConfiguration(ctx, &params)
		return true, reply(ctx, nil, err)
	case "workspace/didChangeWatchedFiles":
		var params DidChangeWatchedFilesParams
		if err := r.UnmarshalParams(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}
		err := server.DidChangeWatchedFiles(ctx, &params)
		return true, reply(ctx, nil, err)
	case "workspace/didChangeWorkspaceFolders":
		var params DidChangeWorkspaceFoldersParams
		if err := r.UnmarshalParams(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}
		err := server.DidChangeWorkspaceFolders(ctx, &params)
		return true, reply(ctx, nil, err)
	case "workspace/didCreateFiles":
		var params CreateFilesParams
		if err := r.UnmarshalParams(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}
		err := server.DidCreateFiles(ctx, &params)
		return true, reply(ctx, nil, err)
	case "workspace/didDeleteFiles":
		var params DeleteFilesParams
		if err := r.UnmarshalParams(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}
		err := server.DidDeleteFiles(ctx, &params)
		return true, reply(ctx, nil, err)
	case "workspace/didRenameFiles":
		var params RenameFilesParams
		if err := r.UnmarshalParams(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}
		err := server.DidRenameFiles(ctx, &params)
		return true, reply(ctx, nil, err)
	case "workspace/executeCommand":
		var params ExecuteCommandParams
		if err := r.UnmarshalParams(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}
		resp, err := server.ExecuteCommand(ctx, &params)
		if err != nil {
			return true, reply(ctx, nil, err)
		}
		return true, reply(ctx, resp, nil)
	case "workspace/symbol":
		var params WorkspaceSymbolParams
		if err := r.UnmarshalParams(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}
		resp, err := server.Symbol(ctx, &params)
		if err != nil {
			return true, reply(ctx, nil, err)
		}
		return true, reply(ctx, resp, nil)
	case "workspace/willCreateFiles":
		var params CreateFilesParams
		if err := r.UnmarshalParams(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}
		resp, err := server.WillCreateFiles(ctx, &params)
		if err != nil {
			return true, reply(ctx, nil, err)
		}
		return true, reply(ctx, resp, nil)
	case "workspace/willDeleteFiles":
		var params DeleteFilesParams
		if err := r.UnmarshalParams(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}
		resp, err := server.WillDeleteFiles(ctx, &params)
		if err != nil {
			return true, reply(ctx, nil, err)
		}
		return true, reply(ctx, resp, nil)
	case "workspace/willRenameFiles":
		var params RenameFilesParams
		if err := r.UnmarshalParams(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}
		resp, err := server.WillRenameFiles(ctx, &params)
		if err != nil {
			return true, reply(ctx, nil, err)
		}
		return true, reply(ctx, resp, nil)
	case "workspaceSymbol/resolve":
		var params WorkspaceSymbol
		if err := r.UnmarshalParams(&params); err != nil {
			return true, sendParseError(ctx, reply, err)
		}
		resp, err := server.ResolveWorkspaceSymbol(ctx, &params)
		if err != nil {
			return true, reply(ctx, nil, err)
		}
		return true, reply(ctx, resp, nil)
	default:
		return false, nil
	}
}

