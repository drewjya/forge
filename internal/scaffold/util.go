package scaffold

import (
	"os"
	"path/filepath"
	"strings"
)

// Reads module name from go.mod
func GoModule() string {
	data, err := os.ReadFile("go.mod")
	if err != nil {
		panic("go.mod not found")
	}
	lines := strings.SplitSeq(string(data), "\n")
	for l := range lines {
		if strings.HasPrefix(l, "module ") {
			return strings.TrimSpace(strings.TrimPrefix(l, "module "))
		}
	}
	panic("module name not found in go.mod")
}
func mkdir(path string) {
	dir := filepath.Clean(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		panic(err)
	}
}

func removeLine(src, line string) string {
	out := []string{}
	for _, l := range strings.Split(src, "\n") {
		if !strings.Contains(l, line) {
			out = append(out, l)
		}
	}
	return strings.Join(out, "\n")
}

func removeBlockContaining(src, marker string) string {
	lines := strings.Split(src, "\n")
	out := []string{}
	skip := false

	for _, l := range lines {
		if strings.Contains(l, marker) {
			skip = true
			continue
		}
		if skip && strings.TrimSpace(l) == "" {
			skip = false
			continue
		}
		if !skip {
			out = append(out, l)
		}
	}

	return strings.Join(out, "\n")
}
