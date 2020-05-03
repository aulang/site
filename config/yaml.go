package config

type Yaml struct {
	Iris struct {
		Port       int    `yaml:"port"`
		Charset    string `yaml:"charset"`
		TimeFormat string `yaml:"time-format"`
	}
	MongoDB struct {
		Uri      string `yaml:"uri"`
		Database string `yaml:"database"`
	}
}
