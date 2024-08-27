package utils

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"sort"
)

// CalculateMD5Hash calculates the MD5 hash of a file's contents
func CalculateMD5Hash(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

// SortFilesByDate sorts the files by oldest first and returns the file path with the oldest file first
func SortFilesByDate(files *[]string) {
	// sort the files by oldest first

	fileInfos := make(map[string][]os.FileInfo, len(*files))
	for _, file := range *files {
		fileInfo, err := os.Stat(file)
		if err != nil {
			continue
		}
		fileInfos[file] = append(fileInfos[file], fileInfo)
	}

	sort.Slice(*files, func(i, j int) bool {
		return fileInfos[(*files)[i]][0].ModTime().Before(fileInfos[(*files)[j]][0].ModTime())
	})
}
