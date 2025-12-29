package scaffold

import (
	"path/filepath"

	"github.com/drewjya/forge/internal/template"
	"github.com/drewjya/forge/internal/util"
)

func AddModule(name string, withCRUD bool) error {
	pkg := util.NewName(name)
	packg := pkg.Package()

	base := filepath.Join("internal/modules", packg)

	mkdirs := []string{
		base + "/" + packg + "controller",
		base + "/" + packg + "service",
		base + "/" + packg + "repository",
		base + "/" + packg + "routes",
	}

	for _, d := range mkdirs {
		mkdir(d)
	}

	// --- always generated ---
	write(base+"/module.go", template.Module(name))
	write(base+"/"+packg+"repository/interface.go", template.RepositoryInterface(name, withCRUD))
	write(base+"/"+packg+"repository/repository.go", template.Repository(name, withCRUD))

	write(base+"/"+packg+"service/interface.go", template.ServiceInterface(name, withCRUD))
	write(base+"/"+packg+"service/service.go", template.Service(name, withCRUD))
	write(base+"/"+packg+"controller/interface.go", template.ControllerInterface(name, withCRUD))
	write(base+"/"+packg+"controller/controller.go", template.Controller(name, withCRUD))

	write(base+"/"+packg+"routes/router.go", template.Router(name, withCRUD))

	// --- CRUD ---
	if withCRUD {
		write(base+"/"+packg+"controller/crud.go", template.ControllerCRUD(name))
		write(base+"/"+packg+"service/crud.go", template.ServiceCRUD(name))
		write(base+"/"+packg+"repository/crud.go", template.RepositoryCRUD(name))
		write(
			"internal/request/"+packg+"_request.go",
			template.Request(name),
		)

	}

	patchBootstrap(packg)
	patchRoutes(name)
	util.GoFmt()
	return nil
}
