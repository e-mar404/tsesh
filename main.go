package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func expandPath(path string) (string, error) {
	if strings.HasPrefix(path, "~") {
		h, err := os.UserHomeDir()
		if err != nil {
			return h, err
		}
		return filepath.Join(h, path[1:]), err
	}
	return path, nil
}

func main() {
	searchPaths := []string{
		"~/projects",
		"~/code",
	}

	for _, path := range searchPaths {
		p, _ := expandPath(path)
		dirEntries, _:= os.ReadDir(p)

		fmt.Printf("directories for %s\n", p)
		for _, entry := range dirEntries {
			fmt.Printf("- %s\n", entry.Name())
		}
	}
}
