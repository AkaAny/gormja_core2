package config

type Config struct {
	Server *ServerConfig
	DB     map[string]DBConfig `mapstructure:"db"`
	Cache  map[string]CacheConfig
	Script map[string]ScriptConfig
}
