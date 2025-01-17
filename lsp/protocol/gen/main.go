package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

var goPlsVersion = flag.String("gopls-version", "v0.17.1", "version of gopls to target")
var out = flag.String("out", "", "directory to write downloaded files to")

// Script to download generated types from gopls sources.
// See https://github.com/golang/tools/blob/master/gopls/internal/lsp/protocol/generate/README.md
func main() {
	flag.Parse()

	source := fmt.Sprintf("https://raw.githubusercontent.com/golang/tools/refs/tags/gopls/%s/gopls/internal/protocol/", *goPlsVersion)
	// source := fmt.Sprintf("https://raw.githubusercontent.com/golang/tools/gopls/%%2F%s/gopls/internal/lsp/protocol/", *goPlsVersion)

	files := []string{"tsjson.go", "tsdocument_changes.go", "tsprotocol.go", "tsclient.go", "tsserver.go", "uri.go"}
	for _, f := range files {
		remote := source + f
		log.Printf("Downloading %s", remote)
		resp, err := http.Get(remote)
		if err != nil {
			panic(err)
		}
		if resp.StatusCode != http.StatusOK {
			panic(fmt.Sprintf("Got %d", resp.StatusCode))
		}
		data, err := io.ReadAll(resp.Body)
		if f == "tsserver.go" || f == "tsclient.go" {
			data = transformJrpc(data)
		} else if f == "uri.go" {
			data = []byte(strings.ReplaceAll(string(data), "golang.org/x/tools/gopls/internal/util/pathutil", "github.com/mightyguava/terragrunt-langserver/lsp/protocol/pathutil"))
		}
		if err != nil {
			panic(err)
		}
		if err := os.WriteFile(path.Join(*out, f), data, 0644); err != nil {
			panic(err)
		}
	}
}

func transformJrpc(data []byte) []byte {
	src := string(data)
	type replace struct {
		from, to string
	}
	for _, r := range []replace{
		{"golang.org/x/tools/internal/jsonrpc2", "github.com/creachadair/jrpc2"},
		{"\"encoding/json\"\n", ""},
		{"r jsonrpc2.Request", "r *jrpc2.Request"},
		{"UnmarshalJSON(r.Params(), &params)", "r.UnmarshalParams(&params)"},
		{"reply jsonrpc2.Replier", "reply Replier"},
	} {
		src = strings.ReplaceAll(src, r.from, r.to)
	}
	return []byte(src)
}
