package configs

type Config struct {
	Version string
}

func (config *Config) Init() {
	config.Version = "0.1.2"
}
