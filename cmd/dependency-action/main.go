package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/codeclysm/extract"
)

func main() {
	if os.Getenv("INPUT_TGZDEPS") == "" {
		os.Exit(0)
	}

	tgzDepURLs := strings.Split(os.Getenv("INPUT_TGZDEPS"), ",")

	for _, tgzDepURL := range tgzDepURLs {
		resp, err := http.Get(tgzDepURL)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: unable to download file at '%s'", tgzDepURL)
			os.Exit(1)
		}
		defer resp.Body.Close()

		if err := extract.Gz(context.Background(), resp.Body, os.Getenv("HOME"), nil); err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: unable to extract file at '%s', ensure it is a valid .tar.gz file", tgzDepURL)
			os.Exit(1)
		}
	}
}
