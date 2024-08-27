package checker

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/amir-the-h/go-duplicate-checker/internal/utils"
)

// CheckDuplicates takes a directory path as an argument, checks for duplicate files by comparing their MD5 hashes,
// and returns a list of duplicates.
func CheckDuplicates(dirPath string) (map[string][]string, error) {
	// Count the total number of files
	fmt.Printf("Scanning directory %s for duplicate files...\n", dirPath)
	fmt.Println("Counting total number of files...")
	totalFiles := 0
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.Mode().IsRegular() {
			totalFiles++
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	fmt.Printf("Total number of files: %d\n", totalFiles)
	bar := utils.ProgressBar(totalFiles, "Scanning files ...")
	fileMap := make(map[string][]string)
	fileMapMutex := sync.Mutex{} // Mutex to synchronize access to fileMap
	signal := make(chan bool)    // make a signal channel to let the application know when the goroutines are done
	err = filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			// run a goroutine to calculate the hash of the file
			go func() {
				fileHash, err := utils.CalculateMD5Hash(path)
				if err != nil {
					signal <- false
					return
				}

				fileMapMutex.Lock()
				fileMap[fileHash] = append(fileMap[fileHash], path)
				fileMapMutex.Unlock()

				bar.Add(1)
				signal <- true
			}()
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	// Wait for all goroutines to finish
	for i := 0; i < totalFiles; i++ {
		<-signal
	}

	bar.Finish()
	fmt.Println()

	bar = utils.ProgressBar(len(fileMap), "Comparing files ...")
	duplicates := make(map[string][]string)
	for hash, files := range fileMap {
		// run a goroutine to compare the files
		go func(hash string, files []string) {
			if len(files) > 1 {
				// sort the files by oldest first
				utils.SortFilesByDate(&files)
				duplicates[hash] = files
			}
			bar.Add(1)
			signal <- true
		}(hash, files)
	}

	// Wait for all goroutines to finish
	for i := 0; i < len(fileMap); i++ {
		<-signal
	}

	bar.Finish()
	fmt.Println()

	return duplicates, nil
}

// PrintDuplicates outputs the list of duplicate files in a human-readable format.
func PrintDuplicates(duplicates map[string][]string) {
	if len(duplicates) == 0 {
		fmt.Println("No duplicate files found.")
		return
	}

	fmt.Println("Duplicate files:")
	for hash, files := range duplicates {
		fmt.Printf("Hash: %s\n", hash)
		for _, file := range files {
			fmt.Printf("\t%s\n", file)
		}
	}
}
