# Terragrunt Language Server

This is a Terragrunt Language server written in Go and a VS Code extension that pairs with it.

## Features

For now, it's fairly barebones in features

* Display HCL syntax errors
* Display Terragrunt file errors
* Cmd+Click to `local` variable declarations
* Cmd+Click to navigate to includes (evaluates Terragrunt functions like `find_in_parent_directories()`)

## Usage

Most users will probably use this as a Visual Studio Code extension. It can be
installed from the [Extension Marketplace](https://marketplace.visualstudio.com/items?itemName=yunchi.terragrunt).

## Structure

```
├── lang
│   ├── terragrunt    -- Terragrunt language support
│   └── testdata      -- Fork of gruntwork-io/terragrunt-infrastructure-live-example for testing
├── lsp
│   ├── document      -- Terragrunt file parsing and evaluation
│   ├── langserver    -- Functionality of the language server, e.g. hover, xrefs
│   └── protocol      -- Implementation of the low level language server protocol
├── main.go           -- Entry point for the server
└── vscode-extension  -- VS Code Extension
```

## Contributing

### Design

The VS Code extension is a thin wrapper around the language server.

The language server is built in Go and uses Terragrunt as a library to support code analysis. The language server
framework is  from `gopls` (the Go language server). When a new `terragrunt.hcl`
is loaded, the language server

* Uses the [HCL v2 parser](https://github.com/hashicorp/hcl) to parse the file into an AST and indexes the AST by file location
* Uses [`terragrunt/config`](https://github.com/gruntwork-io/terragrunt/tree/main/config) as a library to evaluate the Terragrunt file
* Introspections are then performed on the loaded data

This design means that iterating on the extension doesn't require rebuilding Terragrunt / Terraform parsing and evaluation.
That is all already readily available. More of the effort can be directed at building out useful introspections for
the language client.

### Building

This repo uses Hermit to manage binary dependencies (kind of like virtualenv but for all binaries). So source Hermit to start

```shell
source bin/activate-hermit
```

Build and run the language server

```shell
go build -o langserver . && langserver
```

### Testing

```shell
go test ./...
```

For the extension, see [Extension README](./vscode-extension/README.md).
