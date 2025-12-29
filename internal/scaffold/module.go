package scaffold

import (
	"path/filepath"
	"strings"

	"github.com/drewjya/forge/internal/template"
	"github.com/drewjya/forge/internal/util"
)

func AddModule(name string, withCRUD bool) error {
	name = strings.ToLower(name)
	base := filepath.Join("internal/modules", name)

	mkdirs := []string{
		base + "/" + name + "controller",
		base + "/" + name + "service",
		base + "/" + name + "repository",
		base + "/" + name + "routes",
	}

	for _, d := range mkdirs {
		mkdir(d)
	}

	// --- always generated ---
	write(base+"/module.go", template.Module(name))
	write(base+"/"+name+"repository/interface.go", template.RepositoryInterface(name, withCRUD))
	write(base+"/"+name+"repository/repository.go", template.Repository(name, withCRUD))

	write(base+"/"+name+"service/interface.go", template.ServiceInterface(name, withCRUD))
	write(base+"/"+name+"service/service.go", template.Service(name, withCRUD))

	write(base+"/"+name+"controller/interface.go", template.ControllerInterface(name, withCRUD))
	write(base+"/"+name+"controller/controller.go", template.Controller(name, withCRUD))

	write(base+"/"+name+"routes/router.go", template.Router(name, withCRUD))

	// --- CRUD ---
	if withCRUD {
		write(base+"/"+name+"controller/crud.go", template.ControllerCRUD(name))
		write(base+"/"+name+"service/crud.go", template.ServiceCRUD(name))
		write(base+"/"+name+"repository/crud.go", template.RepositoryCRUD(name))
	}

	patchBootstrap(name)
	patchRoutes(name)
	util.GoFmt()
	return nil
}
