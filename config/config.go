package config

import (
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"gormja_core2/cacheutil"
	"gormja_core2/dbutil"
)

type Config struct {
	DB    map[string]DBConfig `mapstructure:"db"`
	Cache map[string]CacheConfig
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

type DBConfig struct {
	ID   string `mapstructure:"id"`
	Type string
	DSN  string `mapstructure:"dsn"`
}

func (x DBConfig) ToDB() *gorm.DB {
	var opener = dbutil.DBOpener{
		Type: x.Type,
		DSN:  x.DSN,
	}
	return opener.ToDB()
}

type CacheConfig struct {
	Type string
	DSN  string `mapstructure:"dsn"`
}

func (x CacheConfig) ToClient() interface{} {
	var opener = cacheutil.CacheOpener{
		Type: x.Type,
		DSN:  x.DSN,
	}
	return opener.ToCache()
}
