package config

type Yaml struct {
	Iris struct {
		Port                int      `yaml:"port"`
		Charset             string   `yaml:"charset"`
		TimeFormat          string   `yaml:"time-format"`
		RemoteAddrHeaders   []string `yaml:"remote-addr-headers"`
		EnableOptimizations bool     `yaml:"enable-optimizations"`
	}
	MongoDB struct {
		Uri      string `yaml:"uri"`
		Database string `yaml:"database"`
	}
}
