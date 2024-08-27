package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/amir-the-h/go-duplicate-checker/internal/checker"
)

func main() {
	dirPath := flag.String("d", "", "Directory path to check for duplicate files")
	remove := flag.Bool("r", false, "Remove duplicate files")
	keepOldest := flag.Bool("o", false, "Keep the oldest version of duplicate files (default: keep newest)")

	flag.Parse()

	if *dirPath == "" {
		*dirPath, _ = os.Getwd()
	}

	duplicates, err := checker.CheckDuplicates(*dirPath)
	if err != nil {
		log.Fatalf("Error checking for duplicates: %v", err)
	}

	if len(duplicates) == 0 {
		fmt.Println("No duplicate files found")
		return
	}

	fmt.Println()

	checker.PrintDuplicates(duplicates)

	if *remove {
		for _, files := range duplicates {
			err = checker.RemoveDuplicates(files, *keepOldest)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
