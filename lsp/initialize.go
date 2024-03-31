package lsp

type InitializeRequest struct {
	Params InitializeRequestParams `json:"params"`
	Request
}

type InitializeRequestParams struct {
	ClientInfo *ClientInfo `json:"clientInfo"`
	// lots more goes here in spec
}

type ClientInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type InitializeResponse struct {
	Response
	Result InitializeResult `json:"result"`
}

type InitializeResult struct {
	ServerInfo   ServerInfo         `json:"serverInfo"`
	Capabilities ServerCapabilities `json:"capabilities"`
}

type ServerCapabilities struct {
	CompletionProvider map[string]any `json:"completionProvider"`
	TextDocumentSync   int            `json:"textDocumentSync"`
	HoverProvider      bool           `json:"hoverProvider"`
	CodeActionProvider bool           `json:"codeActionProvider"`
	DefinitionProvider bool           `json:"definitionProvider"`
}

type ServerInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func NewInitializeResponse(id int) InitializeResponse {
	return InitializeResponse{
		Response: Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: InitializeResult{
			Capabilities: ServerCapabilities{
				TextDocumentSync:   1, // full sync
				HoverProvider:      true,
				DefinitionProvider: true,
				CodeActionProvider: true,
				CompletionProvider: map[string]any{},
			},
			ServerInfo: ServerInfo{
				Name:    "educationallsp",
				Version: "0.0.1",
			},
		},
	}
}
