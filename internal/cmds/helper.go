package cmds

import (
	"os/exec"

	"github.com/maxguuse/bruh/internal/forms"
	"github.com/maxguuse/bruh/internal/types"
	"github.com/samber/lo"
)

func runAskForProjectDetailsForm() (*types.ProjectDetails, error) {
	form := forms.AskForProjectDetails()

	err := form.Run()
	if err != nil {
		return nil, err
	}

	return &types.ProjectDetails{
		Name:  form.GetString("project_name"),
		Owner: form.GetString("project_owner"),
	}, nil
}

func createGoModule(moduleName string, moduleType types.ModuleType, project types.ProjectDetails) error {
	err := createModuleFolder(moduleName, moduleType)
	if err != nil {
		return err
	}

	err = initGoModule(moduleName, moduleType, project)
	if err != nil {
		return err
	}

	err = addGoModuleToWorkspace(moduleName, moduleType)
	if err != nil {
		return err
	}

	return nil
}

func createModuleFolder(moduleName string, moduleType types.ModuleType) error {
	rootDir := lo.If(moduleType == types.App, types.AppsDir).Else(types.LibsDir)

	createAppFolderCmd := exec.Command("mkdir", moduleName)
	createAppFolderCmd.Dir = rootDir
	_, err := createAppFolderCmd.Output()
	if err != nil {
		return err
	}

	return nil
}

func initGoModule(moduleName string, moduleType types.ModuleType, project types.ProjectDetails) error {
	rootDir := lo.If(moduleType == types.App, types.AppsDir).Else(types.LibsDir)

	module := "github.com/" + project.Owner + "/" + project.Name + "/" + rootDir + "/" + moduleName

	goModInitCmd := exec.Command("go", "mod", "init", module)
	goModInitCmd.Dir = rootDir + "/" + moduleName
	_, err := goModInitCmd.Output()
	if err != nil {
		return err
	}

	return nil
}

func addGoModuleToWorkspace(moduleName string, moduleType types.ModuleType) error {
	rootDir := lo.If(moduleType == types.App, types.AppsDir).Else(types.LibsDir)

	addModToWork := exec.Command("go", "work", "use", ".")
	addModToWork.Dir = rootDir + "/" + moduleName
	_, err := addModToWork.Output()
	if err != nil {
		return err
	}

	return nil
}
