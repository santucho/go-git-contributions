package main

import (
	"flag"
	"go-git-contributions/scanner"
	"go-git-contributions/stats"
)

func main() {
	var folder string
	var email string
	flag.StringVar(&folder, "add", "", "add a new folder to scan for Git repositories")
	flag.StringVar(&email, "email", "your@email.com", "the email to scan")
	flag.Parse()

	if folder != "" {
		scanner.Scan(folder)
	}

	stats.Stats(email)
}
