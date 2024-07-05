package main

import (
	"net/http"
	"os"

	"github.com/rajatjindal/wasm-actions-sdk/pkg/fs"
	httpclient "github.com/rajatjindal/wasm-actions-sdk/pkg/http-client"
	"github.com/rajatjindal/wasm-actions-sdk/pkg/ifmt"
)

func main() {
	ifmt.Println("environ", os.Environ())
	response, err := httpclient.Get("https://google.com")
	if err != nil {
		ifmt.Println(err)
		os.Exit(1)
	}

	if response.StatusCode != http.StatusOK {
		ifmt.Printf("expected status code %d, found: %d", http.StatusOK, response.StatusCode)
		// os.Exit(1)
	}

	_ = fs.ListDirs()
}
