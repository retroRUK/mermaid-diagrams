package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/retroRUK/mermaid-diagrams/src/java"
)

func main() {
	// Define flags
	rootPath := flag.String("r", ".", "root directory to scan")
	language := flag.String("l", "", "file ext to parse, ex: java, js")

	// Parse the flags
	flag.Parse()

	if *rootPath == "" {
		flag.PrintDefaults()
		return
	}

	_, err := os.Stat(*rootPath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Printf("root path does not exist: %s\n", *rootPath)
		}
		return
	}

	info, err := os.Stat(*rootPath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Printf("output directory does not exist: %s\n", *rootPath)
		} else {
			fmt.Printf("error checking output path: %v\n", err)
		}
		return
	}

	if !info.IsDir() {
		fmt.Printf("output path must be a directory: %s\n", *rootPath)
		return
	}

	*rootPath = strings.TrimSuffix(*rootPath, "/")

	validLanguages := []string{
		"java",
	}

	if !slices.Contains(validLanguages, *language) {
		flag.PrintDefaults()
		return
	}

	if *language == "java" {
		java.CreateClassDiagrams(*rootPath)
	}
}
