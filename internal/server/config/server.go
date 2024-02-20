package config

import (
	_ "embed"
	"github.com/candbright/go-core/config"
)

//go:embed server.yaml
var ServerData []byte

var ServerConfig *config.Config

func init() {
	cfg, err := config.Parse(ServerData, config.YAML)
	if err != nil {
		panic(err)
	}
	ServerConfig = cfg
}

func Get(key string) string {
	return ServerConfig.Get(key)
}

func GetInt(key string) int {
	return ServerConfig.GetInt(key)
}

func GetInt64(key string) int64 {
	return ServerConfig.GetInt64(key)
}

func GetBool(key string) bool {
	return ServerConfig.GetBool(key)
}
