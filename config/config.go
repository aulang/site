package config

import (
	"io/ioutil"
	"log"
	"strconv"

	"github.com/kataras/iris/v12"
	"gopkg.in/yaml.v2"
)

var Config = new(Yaml)

func init() {
	configFile, err := ioutil.ReadFile("config.yml")

	if err != nil {
		log.Fatalf("加载配置文件失败，%v", err)
	}

	err = yaml.Unmarshal(configFile, Config)

	if err != nil {
		log.Fatalf("解析配置文件失败: %v", err)
	}
}

func Iris() iris.Configurator {
	return iris.WithConfiguration(iris.Configuration{
		Charset:             Config.Iris.Charset,
		TimeFormat:          Config.Iris.TimeFormat,
		RemoteAddrHeaders:   Config.Iris.RemoteAddrHeaders,
		EnableOptimizations: Config.Iris.EnableOptimizations,
	})
}

func Port() string {
	return strconv.Itoa(Config.Iris.Port)
}

const Bucket = "site"
