package lsp

type CodeActionRequest struct {
	Request
	Params TextDocumentCodeActionParams `json:"params"`
}

type TextDocumentCodeActionParams struct {
	context      CodeActionContext      `json:"context"`
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	Range        Range                  `json:"range"`
}

type TextDocumentCodeActionResponse struct {
	Response
	Result []CodeAction `json:"result"`
}

type CodeActionContext struct {
	// add fields as needed
}

type CodeAction struct {
	Edit    *WorkspaceEdit `json:"edit,omitempty"`
	Command *Command       `json:"command,omitempty"`
	Title   string         `json:"title"`
}

type Command struct {
	Title     string        `json:"title"`
	Command   string        `json:"command"`
	Arguments []interface{} `json:"arguments,omitempty"`
}
