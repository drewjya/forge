package scaffold

import (
	"strings"
)

func patchBootstrap(feature string) {
	project := GoModule()
	path := "bootstrap/module.go"

	src := read(path)

	// ---- import ----
	importLine := `"` + project + `/internal/modules/` + feature + `"`

	if !strings.Contains(src, importLine) {
		src = strings.Replace(
			src,
			"// ::forge:imports::",
			importLine+"\n\t// ::forge:imports::",
			1,
		)
	}

	// ---- fx module ----
	moduleLine := feature + ".Module"

	if !strings.Contains(src, moduleLine) {
		src = strings.Replace(
			src,
			"// ::forge:modules::",
			moduleLine+",\n\t// ::forge:modules::",
			1,
		)
	}

	writeRaw(path, src)
}
