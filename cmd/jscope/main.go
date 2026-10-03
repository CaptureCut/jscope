package main

import (
	"fmt"
	"os"

	"jscope/internal/crawler"
	"jscope/internal/fetcher"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: jscope <url>")
		return
	}

	target := os.Args[1]

	html, err := fetcher.Download("https://" + target)
	if err != nil {
		panic(err)
	}

	fmt.Printf("[*] Downloaded %d bytes\n", len(html))

	scripts, err := crawler.GetScriptURLs(html)
	if err != nil {
		panic(err)
	}

	fmt.Printf("[*] Found %d scripts\n", len(scripts))

	for _, script := range scripts {
		fmt.Println(script)
	}
}