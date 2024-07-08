package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/rajatjindal/wasm-actions-sdk/pkg/archive"
	"github.com/rajatjindal/wasm-actions-sdk/pkg/downloads"
)

type Flags struct {
	Version     string
	GitHubToken string
}

func main() {
	flags := &Flags{}

	flagset := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	flagset.StringVar(&flags.Version, "version", "latest", "version to install")
	flagset.StringVar(&flags.GitHubToken, "github-token", "GITHUB_TOKEN", "GitHub token to use for downloading")
	flagset.Parse(os.Args[1:])

	downloadURL := "https://github.com/rajatjindal/kubectl-reverse-proxy/releases/download/v0.3.0/kubectl-reverse-proxy_v0.3.0_darwin_arm64.tar.gz"

	archiveFile, err := downloads.DownloadFile(downloadURL, "download-file")
	if err != nil {
		fmt.Println("error when downloading file ", err)
		return
	}

	err = archive.ExtractTarGzTo(archiveFile, "/tmp/extracted")
	if err != nil {
		fmt.Println(err)
		return
	}

	// find specific version from releases
	// download the archive
	// put it inside tools dir
}
