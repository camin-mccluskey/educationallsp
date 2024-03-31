package main

import (
	"bufio"
	"educationallsp/analysis"
	"educationallsp/lsp"
	"educationallsp/rpc"
	"encoding/json"
	"io"
	"log"
	"os"
)

func main() {
	logger := getLogger("lsp.log")
	logger.Println("Starting the client")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	state := analysis.NewState()
	writer := os.Stdout

	for scanner.Scan() {
		msg := scanner.Bytes()
		method, contents, err := rpc.DecodeMessage(msg)
		if err != nil {
			logger.Printf("Got an error: %s", err)
		}

		handleMessage(logger, writer, state, method, contents)
	}
}

func handleMessage(logger *log.Logger, writer io.Writer, state analysis.State, method string, contents []byte) {
	logger.Printf("Received message with method: %s", method)

	switch method {
	case "initialize":
		var request lsp.InitializeRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("Error unmarshalling initialize request: %s", err)
		}
		logger.Println("Connected to:")
		logger.Printf("ClientInfo Name: %s", request.Params.ClientInfo.Name)
		logger.Printf("ClientInfo Version: %s", request.Params.ClientInfo.Version)

		// respond
		response := lsp.NewInitializeResponse(request.ID)
		writeResponse(writer, response)
	case "textDocument/didOpen":
		var request lsp.DidOpenTextDocumentNotification
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("Error unmarshalling didOpen request: %s", err)
		}
		diagnostics := state.OpenDocument(request.Params.TextDocument.URI, request.Params.TextDocument.Text)
		writeResponse(writer, lsp.PublishDiagnosticsNotification{
			Notification: lsp.Notification{
				RPC:    "2.0",
				Method: "textDocument/publishDiagnostics",
			},
			Params: lsp.PublishDiagnosticsParams{
				Uri:         request.Params.TextDocument.URI,
				Diagnostics: diagnostics,
			},
		})
	case "textDocument/didChange":
		var request lsp.DidChangeTextDocumentNotification
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("Error unmarshalling didChange request: %s", err)
		}
		for _, change := range request.Params.ContentChanges {
			diagnostics := state.UpdateDocument(request.Params.TextDocument.URI, change.Text)
			writeResponse(writer, lsp.PublishDiagnosticsNotification{
				Notification: lsp.Notification{
					RPC:    "2.0",
					Method: "textDocument/publishDiagnostics",
				},
				Params: lsp.PublishDiagnosticsParams{
					Uri:         request.Params.TextDocument.URI,
					Diagnostics: diagnostics,
				},
			})
		}
	case "textDocument/hover":
		var request lsp.HoverRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("Error unmarshalling hover request: %s", err)
			return
		}
		response := state.Hover(request.ID, request.Params.TextDocument.URI, request.Params.Position)
		writeResponse(writer, response)
	case "textDocument/definition":
		var request lsp.DefinitionRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("Error unmarshalling definition request: %s", err)
			return
		}
		// response := state.Hover(request.ID, request.Params.TextDocument.URI, request.Params.Position)
		response := state.Definition(request.ID, request.Params.TextDocument.URI, request.Params.Position)
		writeResponse(writer, response)
	case "textDocument/codeAction":
		var request lsp.CodeActionRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("Error unmarshalling codeAction request: %s", err)
			return
		}
		response := state.TextDocumentCodeAction(request.ID, request.Params.TextDocument.URI)
		writeResponse(writer, response)
	case "textDocument/completion":
		var request lsp.CompletionRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("Error unmarshalling completion request: %s", err)
			return
		}
		response := state.TextDocumentCompletion(request.ID, request.Params.TextDocument.URI)
		writeResponse(writer, response)
	}
}

func writeResponse(writer io.Writer, msg any) {
	reply := rpc.EncodeMessage(msg)
	writer.Write([]byte(reply))
}

func getLogger(filename string) *log.Logger {
	logFile, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		panic("Failed to open log file")
	}
	return log.New(logFile, "[educationallsp] ", log.Ldate|log.LUTC|log.Lshortfile)
}
