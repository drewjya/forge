package util

import (
	"strings"

	"github.com/gertd/go-pluralize"
)

var pluralizer = pluralize.NewClient()

// Input: "happy_hour"
type Name struct {
	raw string
}

// NewName creates a normalized name helper
func NewName(input string) Name {
	return Name{
		raw: strings.ToLower(input),
	}
}

// Snake returns the original snake_case
func (n Name) Snake() string {
	return n.raw
}

// Package returns a Go-safe package name (no underscores)
func (n Name) Package() string {
	return strings.ReplaceAll(n.raw, "_", "")
}

// Singular snake_case
func (n Name) Singular() string {
	return pluralizer.Singular(n.raw)
}

// Plural snake_case
func (n Name) Plural() string {
	return pluralizer.Plural(n.raw)
}

// Title returns Go-style struct name
// happy_hour -> HappyHour
func (n Name) Title() string {
	parts := strings.Split(n.Singular(), "_")
	for i, p := range parts {
		parts[i] = Title(p)
	}
	return strings.Join(parts, "")
}

// TitlePlural returns plural Go-style struct name
// happy_hour -> HappyHours
func (n Name) TitlePlural() string {
	parts := strings.Split(n.Plural(), "_")
	for i, p := range parts {
		parts[i] = Title(p)
	}
	return strings.Join(parts, "")
}

// Kebab returns kebab-case for routes
// happy_hour -> happy-hour
func (n Name) Kebab() string {
	return kebab(n.raw)
}

func kebab(n string) string {
	return strings.ReplaceAll(n, "_", "-")
}

func (n Name) KebabPlural() string {
	println(n.raw)
	parts := strings.Split(n.Plural(), "_")
	for i, p := range parts {
		parts[i] = kebab(p)
	}
	return strings.Join(parts, "-")
}
