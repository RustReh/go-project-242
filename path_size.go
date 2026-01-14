package code

import (
	"fmt"
	"os"
)

func GetPathSize(path string, recursive, human, all bool) (string, error) {
	info, err := os.Lstat(path)
	if err != nil {
		return "", err
	}

	if !info.IsDir() {
		return fmt.Sprintf("%d\t%s", info.Size(), path), nil
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		return "", err
	}

	totalSize := int64(0)
	for _, entry := range entries {
		entryInfo, err := entry.Info()
		if err != nil {
			continue
		}
		totalSize += entryInfo.Size()
	}

	return fmt.Sprintf("%d\t%s", totalSize, path), nil
}
