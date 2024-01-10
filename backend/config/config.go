package config

type Config struct {
	Server      *ServerConfig
	DB          map[string]DBConfig `mapstructure:"db"`
	EnableCache bool
	Cache       map[string]CacheConfig
	Script      map[string]ScriptConfig
}
