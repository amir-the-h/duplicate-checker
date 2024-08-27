package checker

import (
	"fmt"
	"os"

	"github.com/amir-the-h/go-duplicate-checker/internal/utils"
)

// RemoveDuplicates removes duplicate files based on the specified option (keep newest or oldest).
func RemoveDuplicates(duplicates []string, keepOldest bool) error {
	if len(duplicates) == 0 {
		return nil
	}

	// Create a new progress bar for removing files
	bar := utils.ProgressBar(len(duplicates)-1, "Removing files ...")

	// Remove the duplicates based on the specified option
	if keepOldest {
		// remove the first file from the duplicates array
		duplicates = duplicates[1:]
	} else {
		// remove the last file from the duplicates array
		duplicates = duplicates[:len(duplicates)-1]
	}

	for i := 0; i < len(duplicates); i++ {
		err := os.Remove(duplicates[i])
		if err != nil {
			return fmt.Errorf("error removing file %s: %v", duplicates[i], err)
		}

		bar.Add(1)
	}

	bar.Finish()

	return nil
}
