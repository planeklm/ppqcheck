package main

import (
	"flag"
	"fmt"
	"os"

	"go.mozilla.org/pkcs7"
	"howett.net/plist"
)

// HasPPQ reports whether the mobileprovision has `PPQCheck` set to `true`.
func HasPPQ(path string) (bool, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return false, fmt.Errorf("failed to read file: %w", err)
	}

	p7, err := pkcs7.Parse(data)
	if err != nil {
		return false, fmt.Errorf("failed to parse cms data: %w", err)
	}

	var provision map[string]any
	if _, err = plist.Unmarshal(p7.Content, &provision); err != nil {
		return false, fmt.Errorf("failed to parse plist: %w", err)
	}

	ppq, ok := provision["PPQCheck"]
	return ok && ppq.(bool), nil // this is fine since it checking `ok` will short-circuit
}

func main() {
	path := flag.String("path", "", "Path to the .mobileprovision (required)")
	flag.Parse()

	if *path == "" {
		fmt.Println("error: --path flag is needed")
		flag.Usage()
		os.Exit(1)
	}

	ppq, err := HasPPQ(*path)
	if err != nil {
		fmt.Printf("error checking for ppq: %v", err)
		os.Exit(1)
	}

	fmt.Printf("PPQ detected: %t\n", ppq)
}
