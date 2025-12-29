package scaffold

import (
	"strings"

	"github.com/andre/forge/internal/util"
)

func patchRoutes(feature string) {
	project := GoModule()
	path := "internal/routes/routes.go"

	src := read(path)

	Feature := util.Title(feature)

	// ---- imports ----
	imports := []string{
		`"` + project + `/internal/modules/` + feature + `/` + feature + `controller"`,
		`"` + project + `/internal/modules/` + feature + `/` + feature + `routes"`,
	}

	for _, imp := range imports {
		if !strings.Contains(src, imp) {
			src = strings.Replace(
				src,
				"// ::forge:imports::",
				imp+"\n\t// ::forge:imports::",
				1,
			)
		}
	}

	// ---- deps ----
	depLine := Feature + "Handler " +
		feature + "controller.I" + Feature + "Controller"

	if !strings.Contains(src, depLine) {
		src = strings.Replace(
			src,
			"// ::forge:deps::",
			"\t"+depLine+"\n\t// ::forge:deps::",
			1,
		)
	}

	// ---- routes ----
	routeBlock := `
	` + feature + `Router := ` + feature + `routes.New` + Feature + `Router(deps.` + Feature + `Handler)
	` + feature + `Router.RegisterRoutes(api)
`

	if !strings.Contains(src, feature+"Router :=") {
		src = strings.Replace(
			src,
			"// ::forge:routes::",
			routeBlock+"\n\t// ::forge:routes::",
			1,
		)
	}

	writeRaw(path, src)
}
