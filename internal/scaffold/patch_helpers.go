package scaffold

import (
	"strings"
)

// Inserts an import into a Go import block
func insertImport(src string, importLine string) string {
	if strings.Contains(src, importLine) {
		return src
	}

	return strings.Replace(
		src,
		"import (",
		"import (\n\t"+importLine,
		1,
	)
}

// Inserts a field into a struct definition
func insertIntoStruct(src, structName, fieldLine string) string {
	marker := "type " + structName + " struct {"
	if strings.Contains(src, fieldLine) {
		return src
	}

	return strings.Replace(
		src,
		marker,
		marker+"\n\t"+fieldLine,
		1,
	)
}
