package lsp

type DidChangeTextDocumentNotification struct {
	Notification
	Params TextDocumentDidChangeTextDocumentParams `json:"params"`
}

type TextDocumentDidChangeTextDocumentParams struct {
	TextDocument   VersionedTextDocumentIdentifier `json:"textDocument"`
	ContentChanges []TextDocumentContentChangeEven `json:"contentChanges"`
}

type TextDocumentContentChangeEven struct {
	Text string `json:"text"`
}
