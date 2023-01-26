package config

import "gormja_core2/cacheutil"

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
