package analysis

import (
	"educationallsp/lsp"
	"fmt"
	"strings"
)

type State struct {
	Documents map[string]string // map of filenames to contents
}

func NewState() State {
	return State{Documents: map[string]string{}}
}

func getDiagnosticsForFile(text string) []lsp.Diagnostic {
	diagnostics := []lsp.Diagnostic{}
	badStr := "VS Code"
	for row, line := range strings.Split(text, "\n") {
		if strings.Contains(line, badStr) {
			idx := strings.Index(line, badStr)
			diagnostics = append(diagnostics, lsp.Diagnostic{
				Range:    LineRange(row, idx, idx+len(badStr)),
				Severity: 1, Source: "Common sense",
				Message: "VS Code is bad, use Neovim instead",
			})
		}
	}
	return diagnostics
}

func (s *State) OpenDocument(uri, text string) []lsp.Diagnostic {
	s.Documents[uri] = text
	return getDiagnosticsForFile(text)
}

func (s *State) UpdateDocument(uri, text string) []lsp.Diagnostic {
	s.Documents[uri] = text
	return getDiagnosticsForFile(text)
}

func (s *State) Hover(id int, uri string, position lsp.Position) lsp.HoverResponse {
	// irl this would actually do something useful to send a message back given the hovered text
	document := s.Documents[uri]

	return lsp.HoverResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: lsp.HoverResult{
			Contents: fmt.Sprintf("File: %s, Characters: %d", uri, len(document)),
		},
	}
}

func (s *State) Definition(id int, uri string, position lsp.Position) lsp.DefinitionResponse {
	return lsp.DefinitionResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: lsp.Location{
			URI: uri,
			Range: lsp.Range{
				Start: lsp.Position{
					Line:      position.Line - 1,
					Character: 0,
				},
				End: lsp.Position{
					Line:      position.Line,
					Character: 0,
				},
			},
		},
	}
}

func (s *State) TextDocumentCodeAction(id int, uri string) lsp.TextDocumentCodeActionResponse {
	text := s.Documents[uri]

	actions := []lsp.CodeAction{}
	badStr := "VS Code"
	for row, line := range strings.Split(text, "\n") {
		idx := strings.Index(line, badStr)
		if idx >= 0 {
			replaceChange := map[string][]lsp.TextEdit{}
			replaceChange[uri] = []lsp.TextEdit{
				{Range: LineRange(row, idx, idx+len(badStr)), NewText: "Neovim"},
			}
			actions = append(actions, lsp.CodeAction{
				Title: "Replace VS Code with Neovim",
				Edit:  &lsp.WorkspaceEdit{Changes: replaceChange},
			})
			censorChange := map[string][]lsp.TextEdit{}
			censorChange[uri] = []lsp.TextEdit{
				{Range: LineRange(row, idx, idx+len(badStr)), NewText: "VS C*de"},
			}
			actions = append(actions, lsp.CodeAction{
				Title: "Censor VS Code",
				Edit:  &lsp.WorkspaceEdit{Changes: censorChange},
			})
		}
	}

	response := lsp.TextDocumentCodeActionResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: actions,
	}
	return response
}

func LineRange(line, start, end int) lsp.Range {
	return lsp.Range{
		Start: lsp.Position{
			Line:      line,
			Character: start,
		},
		End: lsp.Position{
			Line:      line,
			Character: end,
		},
	}
}

func (s *State) TextDocumentCompletion(id int, uri string) lsp.CompletionResponse {
	// Ask your static analysis tools to figure out good completions
	// we're just hard coding a single completion result here
	items := []lsp.CompletionItem{
		{
			Label:         "Neovim",
			Detail:        "Very good editor",
			Documentation: "Fun to use",
		},
	}

	return lsp.CompletionResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: items,
	}
}
