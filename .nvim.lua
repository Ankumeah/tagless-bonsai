local lsp = "gopls"
local capabilities = require('cmp_nvim_lsp').default_capabilities()
local opts = { capabilities = capabilities }

opts.settings = {
  gopls = {
    env = {
      GOOS = "js",
      GOARCH = "wasm"
    },
    buildFlags = { "-tags=js,wasm" }
  }
}

vim.lsp.config(lsp, opts)
for _, client in ipairs(vim.lsp.get_clients({ name = "gopls" })) do
  client:stop({ force = true })
end
