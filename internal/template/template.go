package template

import (
	"os"
	"strings"
)

func goModule() string {
	data, err := os.ReadFile("go.mod")
	if err != nil {
		panic("go.mod not found (run forge inside a Go module)")
	}

	lines := strings.SplitSeq(string(data), "\n")
	for line := range lines {
		line = strings.TrimSpace(line)
		if after, ok :=strings.CutPrefix(line, "module "); ok  {
			return strings.TrimSpace(after)
		}
	}

	panic("module name not found in go.mod")
}
