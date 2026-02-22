package picker

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	"github.com/e-mar404/tsesh/config"
)

func expandPath(path string) string {
	expanded := path
	if strings.Contains(path, "~") {
		home, _ := os.UserHomeDir()
		expanded = filepath.Join(home, path[1:])
	}
	return expanded
}

func searchPaths(cfg config.Search) []list.Item {
	m := make(map[string]Item)
	dirList := []list.Item{}
	for _, root := range cfg.Paths {
		expandedRoot := expandPath(root)
		filepath.WalkDir(expandedRoot, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				fmt.Fprintf(os.Stderr, "error while walking dir: %v\n", err)
				return err
			}

			if !d.IsDir() {
				return nil
			}

			sessionName := d.Name()
			// paths that are explicitly on the search path list always get searched
			if path != expandedRoot {
				pattern := cfg.IgnorePattern
				if cfg.IgnorePattern == "" {
					pattern = "![*]"
				}

				match, err := regexp.MatchString(pattern, sessionName)
				if err != nil {
					return err
				}

				if match {
					return filepath.SkipDir
				}

				if strings.HasPrefix(sessionName, ".") && cfg.IgnoreHidden {
					return filepath.SkipDir
				}
			}

			if strings.Contains(d.Name(), ".") {
				sessionName = "_" + sessionName[1:]
			}

			if ok := m[sessionName]; ok == (Item{}) {
				item := Item{
					SessionName: sessionName,
					Path:        path,
				}
				dirList = append(dirList, item)
				m[sessionName] = item
			}

			// placing check after maintains desired max depth of 1 behavior
			if d.IsDir() && path != expandedRoot {
				return filepath.SkipDir
			}

			return nil
		})
	}
	return dirList
}
