package code

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)


func GetPathSize(path string, recursive, human, all bool) (string, error) {
	info, err := os.Lstat(path)
	if err != nil {
		return "", err
	}
	if !info.IsDir() {
		fileSize := info.Size()
		sizeFormated := FormatSize(fileSize, human)
		return fmt.Sprintf("%s %s", sizeFormated, path), nil
	}

	totalSize, err := calculateDirSize(path, recursive, all)
	if err != nil {
		return "", err
	}

	sizeFormatted := FormatSize(totalSize, human)
	return fmt.Sprintf("%s %s", sizeFormatted, path), nil
}

func calculateDirSize(path string, recursive, all bool) (int64, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return 0, err
	}

	totalSize := int64(0)
	for _, entry := range entries {
		if !all && strings.HasPrefix(entry.Name(), ".") {
			continue
		}

		entryInfo, err := entry.Info()
		if err != nil {
			continue
		}

	if entry.IsDir() && recursive {
		subPath := filepath.Join(path, entry.Name())
		subSize, err := calculateDirSize(subPath, recursive, all)
		if err != nil {
			continue
		}
		totalSize += subSize
	} else if !entry.IsDir() {
		totalSize += entryInfo.Size()
	}
	}

	return totalSize, nil
}


func FormatSize(size int64, human bool) string {
	if !human {
		return fmt.Sprintf("%dB", size)
	}

	const (
		KB = 1024
		MB = 1024 * 1024
		GB = 1024 * 1024 * 1024
		TB = 1024 * 1024 * 1024 * 1024
		PB = 1024 * 1024 * 1024 * 1024 * 1024
		EB = 1024 * 1024 * 1024 * 1024 * 1024 * 1024
	)

	switch {
	case size < KB:
		return fmt.Sprintf("%dB", size)
	case size < MB:
		return fmt.Sprintf("%.1fKB", float64(size)/float64(KB))
	case size < GB:
		return fmt.Sprintf("%.1fMB", float64(size)/float64(MB))
	case size < TB:
		return fmt.Sprintf("%.1fGB", float64(size)/float64(GB))
	case size < PB:
		return fmt.Sprintf("%.1fTB", float64(size)/float64(TB))
	case size < EB:
		return fmt.Sprintf("%.1fPB", float64(size)/float64(PB))
	default:
		return fmt.Sprintf("%.1fEB", float64(size)/float64(EB))
	}
}
