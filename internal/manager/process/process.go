package process

type Config struct {
	RootDir string
}

type Process struct {
	AllowList        *AllowList
	ServerProperties *ServerProperties
}

func New(cfg *Config) *Process {
	manager := &Process{
		AllowList:        NewAllowList(cfg.RootDir),
		ServerProperties: NewServerProperties(cfg.RootDir),
	}
	return manager
}
