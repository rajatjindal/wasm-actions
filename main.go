package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/rajatjindal/wasm-actions-sdk/pkg/archive"
	"github.com/rajatjindal/wasm-actions-sdk/pkg/downloads"
	"github.com/rajatjindal/wasm-actions-sdk/pkg/platform"
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

	// for some reason spin up --flag-name does not work
	// workaround is spin up - --flag-name
	// and yes, spin up -- --flag-name does not work either
	if len(os.Args) >= 3 {
		flagset.Parse(os.Args[2:])
	}

	dir, err := os.MkdirTemp("/spin-ghactions-tmp", "")
	if err != nil {
		fmt.Println(err)
		return
	}
	// tinygo dont support it yet
	defer os.RemoveAll(dir)

	downloadFile := filepath.Join(dir, "download-file")
	runtimeOS, runtimeArch, err := platform.GetPlatform()
	if err != nil {
		fmt.Println(err)
		return
	}

	downloadURL := fmt.Sprintf("https://github.com/rajatjindal/kubectl-reverse-proxy/releases/download/%s/kubectl-reverse-proxy_%s_%s_%s.tar.gz", flags.Version, flags.Version, runtimeOS, runtimeArch)
	archiveFile, err := downloads.DownloadFile(downloadURL, downloadFile)
	if err != nil {
		fmt.Println("error when downloading file ", err)
		return
	}

	err = archive.ExtractTarGzTo(archiveFile, "/spin-ghactions-tmp")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = os.Chmod("/spin-ghactions-tmp/kubectl-reverse_proxy", 0755)
	if err != nil {
		fmt.Println(err)
		return
	}

	// find specific version from releases
	// download the archive. ensure trigger pass underlying os/platform
	// put it inside tools dir
}
