package scaffold

import "github.com/drewjya/forge/internal/util"

func removeFromBootstrap(feature string) {
	project := GoModule()
	path := "bootstrap/module.go"
	src := read(path)
	nn := util.NewName(feature)
	feature = nn.Package()

	importLine := `"` + project + `/internal/modules/` + feature + `"`
	moduleLine := feature + ".Module,"

	src = removeLine(src, importLine)
	src = removeLine(src, moduleLine)

	writeRaw(path, src)
}
