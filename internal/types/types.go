package types

type Config struct {
	Project ProjectDetails `yaml:"project"`
}

type ProjectDetails struct {
	Name  string `yaml:"name"`
	Owner string `yaml:"owner"`
}

type ModuleType int

const (
	App ModuleType = iota
	Lib
)

type SubcommandType int

const (
	InitProject SubcommandType = iota
	CreateModule
)

const (
	AppsDir = "apps"
	LibsDir = "libs"
)
