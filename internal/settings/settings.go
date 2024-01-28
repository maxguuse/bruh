package settings

type Settings struct {
	Project *ProjectDetails `yaml:"project"`
}

type ProjectDetails struct {
	Name  string `yaml:"name"`
	Owner string `yaml:"owner"`
}
