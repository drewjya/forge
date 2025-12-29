package scaffold

import (
	"os"
	"path/filepath"

	"github.com/andre/forge/internal/template"
	"github.com/andre/forge/internal/util"
)

func InitProject(module string) error {
	dirs := []string{
		"bootstrap/config",
		"bootstrap/db",
		"bootstrap/log", // 👈 ensure dir
		"bootstrap/server",
		"bootstrap/validation",
		"cmd/api",
		"internal/modules",
		"internal/routes",
		"internal/types",
		"migrations",
	}

	for _, d := range dirs {
		if err := os.MkdirAll(filepath.Join(module, d), 0755); err != nil {
			return err
		}
	}

	write(module+"/go.mod", template.GoMod(module))
	write(module+"/.env", template.Env())
	write(module+"/bobgen.yaml", template.BobgenYAML())

	write(module+"/bootstrap/module.go", template.BootstrapRoot(module))
	write(module+"/bootstrap/config/config.go", template.Config())
	write(module+"/bootstrap/config/module.go", template.ConfigModule())

	write(module+"/bootstrap/db/postgres.go", template.DB(module))

	write(module+"/bootstrap/log/log.go", template.Log())
	write(module+"/bootstrap/log/module.go", template.LogModule())

	write(module+"/bootstrap/server/server.go", template.Server(module))

	write(module+"/internal/routes/module.go", template.RoutesModule())
	write(module+"/internal/routes/routes.go", template.Routes())

	write(module+"/cmd/api/main.go", template.Main(module))

	// ---------- migrations ----------
	write(module+"/migrations/001_extensions.up.sql", template.MigrationExtensionsUp())
	write(module+"/migrations/001_extensions.down.sql", template.MigrationExtensionsDown())

	util.GoFmt()
	return nil
}
