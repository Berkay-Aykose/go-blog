package helpers

import (
	"path/filepath"
)

func Include(path string) []string {
	files, _ := filepath.Glob("site/views/templates/*.html")
	path_files, _ := filepath.Glob("site/views/" + path + "/*.html")

	for i := 0; i < len(path_files); i++ {
		files = append(files, path_files[i])
	}

	return files
}
