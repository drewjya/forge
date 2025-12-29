package scaffold

import (
	"github.com/drewjya/forge/internal/util"
)

func removeFromRoutes(feature string) {
	project := GoModule()
	nn := util.NewName(feature)
	feature = nn.Package()
	Feature := nn.Title()
	println(nn.Snake() + " jsdjsdj")
	path := "internal/routes/routes.go"
	src := read(path)

	imports := []string{
		`"` + project + `/internal/modules/` + feature + `/` + feature + `controller"`,
		`"` + project + `/internal/modules/` + feature + `/` + feature + `routes"`,
	}

	for _, imp := range imports {
		src = removeLine(src, imp)
	}

	depLine := Feature + "Handler " +
		feature + "controller.I" + Feature + "Controller"
	src = removeLine(src, depLine)

	routerBlock := feature + "Router :="
	src = removeBlockContaining(src, routerBlock)

	writeRaw(path, src)
}
