package config

import (
	"gorm.io/gorm"
	"gormja_core2/dbutil"
)

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
