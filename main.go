package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	path := flag.String("path", ".", "path to git repository")
	flag.Parse()

	stats, err := CollectStats(*path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	profile := GenerateProfile(stats)
	PrintProfile(profile, stats)
