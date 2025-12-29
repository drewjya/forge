package scaffold

import (
	"os"
	"path/filepath"
)

// write creates parent dirs automatically
func write(path string, content string) {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		panic(err)
	}

	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		panic(err)
	}
}

// writeRaw overwrites file without mkdir logic
func writeRaw(path string, content string) {
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		panic(err)
	}
}

// read reads file or panics (scaffold-time error)
func read(path string) string {
	b, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(b)
}
