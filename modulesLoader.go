package abstractions

type IModulesLoader interface {
	Load(path string) error
	LoadModules() error
	GetLoadedModules() []string
}
