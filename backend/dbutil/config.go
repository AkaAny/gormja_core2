package dbutil

import (
	"fmt"
	gorm_oracle "github.com/mudita33/gorm-ora"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBOpener struct {
	Type string
	DSN  string
}

type DialectorOpenFunc = func(dsn string) gorm.Dialector
type GormOpenFunc = func(dialector gorm.Dialector, opts ...gorm.Option) (*gorm.DB, error)

func (dbc DBOpener) ToDB() *gorm.DB {
	var typeDialectOpenMap = map[string]DialectorOpenFunc{
		"mysql":  mysql.Open,
		"oracle": gorm_oracle.Open,
	}
	fmt.Println(dbc.DSN)
	var dialector = typeDialectOpenMap[dbc.Type](dbc.DSN)
	var typeGormOpenMap = map[string]GormOpenFunc{
		"mysql":  gorm.Open,
		"oracle": OpenOracleWithWatcher,
	}
	db, err := typeGormOpenMap[dbc.Type](dialector)
	if err != nil {
		panic(err)
	}
	return db
}
