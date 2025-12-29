package scaffold

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/andre/forge/internal/util"
)

func RemoveModule(feature string) error {
	feature = strings.ToLower(feature)

	// 1. delete module folder
	modPath := filepath.Join("internal/modules", feature)
	if err := os.RemoveAll(modPath); err != nil {
		return err
	}

	// 2. patch bootstrap + routes
	removeFromBootstrap(feature)
	removeFromRoutes(feature)
	util.GoFmt()
	return nil
}
