package types

type Module struct {
	Name string
	Type ModuleType
}

type ModuleType int

const (
	App ModuleType = iota
	Lib
)

func (m *ModuleType) String() string {
	return [...]string{
		AppsDir,
		LibsDir,
	}[*m]
}

type SubcommandType int

const (
	InitProject SubcommandType = iota
	CreateModule
)

const (
	AppsDir = "apps"
	LibsDir = "libs"
)
