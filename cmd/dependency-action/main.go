package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/codeclysm/extract"
)

func main() {
	fmt.Println("oh boy, here I go downloadin' again!")

	tgzDepURL := os.Getenv("INPUT_TGZDEPS")
	if tgzDepURL == "" {
		os.Exit(0)
	}

	resp, err := http.Get(tgzDepURL)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if err := extract.Gz(context.Background(), resp.Body, os.Getenv("HOME"), nil); err != nil {
		log.Fatal(err)
	}
}
