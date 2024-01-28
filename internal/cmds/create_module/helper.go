package create_module

import (
	"fmt"

	"github.com/maxguuse/bruh/internal/fs"
	"github.com/maxguuse/bruh/internal/settings"
	"github.com/maxguuse/bruh/internal/types"
)

func createGoModule(module *types.Module, project *settings.ProjectDetails) error {
	moduleTypeString := module.Type.String()
	root := moduleTypeString + "/" + module.Name

	err := fs.Mkdir(moduleTypeString, module.Name)
	if err != nil {
		return err
	}

	importPath := fmt.Sprintf(
		"github.com/%s/%s/%s/%s",
		project.Owner, project.Name, moduleTypeString, module.Name,
	)

	err = fs.GoModInit(root, importPath)
	if err != nil {
		return err
	}

	err = fs.GoWorkUse(module)
	if err != nil {
		return err
	}

	return nil
}
