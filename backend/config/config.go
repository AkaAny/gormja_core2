package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Server *ServerConfig
	DB     map[string]DBConfig `mapstructure:"db"`
	Cache  map[string]CacheConfig
	Script map[string]ScriptConfig
}

func Unmarshal(vh *viper.Viper) *Config {
	var cfg = new(Config)
	if err := vh.Unmarshal(cfg); err != nil {
		panic(err)
	}
	return cfg
}

func NewDefaultViper() *viper.Viper {
	var vh = viper.New()
	vh.AddConfigPath("config_gormja_core2")
	vh.SetConfigName("conf")
	if err := vh.ReadInConfig(); err != nil {
		panic(err)
	}
	return vh
}
