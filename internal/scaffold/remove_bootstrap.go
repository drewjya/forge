package scaffold

func removeFromBootstrap(feature string) {
	project := GoModule()
	path := "bootstrap/module.go"
	src := read(path)

	importLine := `"` + project + `/internal/modules/` + feature + `"`
	moduleLine := feature + ".Module,"

	src = removeLine(src, importLine)
	src = removeLine(src, moduleLine)

	writeRaw(path, src)
}
