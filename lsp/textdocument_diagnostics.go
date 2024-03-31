package lsp

type PublishDiagnosticsNotification struct {
	Notification
	Params PublishDiagnosticsParams `json:"params"`
}

type PublishDiagnosticsParams struct {
	Uri         string       `json:"uri"`
	Diagnostics []Diagnostic `json:"diagnostics"`
}

type Diagnostic struct {
	Source   string `json:"source"`
	Message  string `json:"message"`
	Range    Range  `json:"range"`
	Severity int    `json:"severity"`
}
