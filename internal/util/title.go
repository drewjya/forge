package util

import (
	"os"
	"os/exec"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func Title(s string) string {
	caser := cases.Title(language.Und)
	return caser.String(strings.ToLower(strings.TrimSpace(s)))
}

func GoFmt() {
	cmd := exec.Command("gofmt", "-s", "-w", ".")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		panic("gofmt failed")
	}
}
