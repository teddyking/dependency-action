package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/codeclysm/extract"
)

func main() {
	fmt.Println("oh boy, here I go downloadin' again!")

	if os.Getenv("INPUT_TGZDEPS") == "" {
		os.Exit(0)
	}

	tgzDepURLs := strings.Split(os.Getenv("INPUT_TGZDEPS"), ",")

	for _, tgzDepURL := range tgzDepURLs {
		resp, err := http.Get(tgzDepURL)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		if err := extract.Gz(context.Background(), resp.Body, os.Getenv("HOME"), nil); err != nil {
			log.Fatal(err)
		}
	}
}
