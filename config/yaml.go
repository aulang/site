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
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	}

	Minio struct {
		Endpoint  string `yaml:"endpoint"`
		AccessKey string `yaml:"accessKey"`
		SecretKey string `yaml:"secretKey"`
		UseSSL    bool   `yaml:"useSSL"`
	}
}
