# Educational LSP

This is an example LSP server implementation. Based on the [video](https://www.youtube.com/watch?v=YsdlcQoHqPY&t=3034s) by TJ DeVries.

The server is a very incomplete implementation to the [Language Server Protocol Specification](https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/).

## Installation & Testing

Build the server:

```sh
go build main.go -o bin
```

Make your editor aware of the LSP server (Neovim example, but should work with VS Code etc...):

```lua
local client = vim.lsp.start_client({
  name = "educationallsp",
  cmd = {
    "<path-to-server-executable>" -- "/Users/caminmccluskey/projects/educationallsp/bin/educationallsp"
  },
})

if not client then
  vim.notify("No LSP found")
  return
end

vim.api.nvim_create_autocmd("FileType", {
  pattern = "markdown", --  or any file type frankly, there is nothing specific to Markdown handled by this LSP server
  callback = function()
    vim.lsp.buf_attach_client(0, client)
  end,
})
```
