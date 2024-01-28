package cmds

import (
	"os/exec"

	"github.com/maxguuse/bruh/internal/types"
)

func createGoModule(module *types.Module, project types.ProjectDetails) error {
	err := createModuleFolder(module)
	if err != nil {
		return err
	}

	err = initGoModule(module, project)
	if err != nil {
		return err
	}

	err = addGoModuleToWorkspace(module)
	if err != nil {
		return err
	}

	return nil
}

func createModuleFolder(module *types.Module) error {
	createAppFolderCmd := exec.Command("mkdir", module.Name)
	createAppFolderCmd.Dir = module.Type.String()
	_, err := createAppFolderCmd.Output()
	if err != nil {
		return err
	}

	return nil
}

func initGoModule(module *types.Module, project types.ProjectDetails) error {
	rootDir := module.Type.String()

	importPath := "github.com/" + project.Owner + "/" + project.Name + "/" + rootDir + "/" + module.Name

	goModInitCmd := exec.Command("go", "mod", "init", importPath)
	goModInitCmd.Dir = rootDir + "/" + module.Name
	_, err := goModInitCmd.Output()
	if err != nil {
		return err
	}

	return nil
}

func addGoModuleToWorkspace(module *types.Module) error {
	rootDir := module.Type.String()

	addModToWork := exec.Command("go", "work", "use", ".")
	addModToWork.Dir = rootDir + "/" + module.Name
	_, err := addModToWork.Output()
	if err != nil {
		return err
	}

	return nil
}
