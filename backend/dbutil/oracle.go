package dbutil

import (
	"database/sql"
	"fmt"
	gorm_oracle "github.com/mudita33/gorm-ora"
	"gorm.io/gorm"
	"strings"
)

type Banner struct {
	Banner string `gorm:"column:BANNER"`
}

func OpenOracleWithWatcher(dialector gorm.Dialector, opts ...gorm.Option) (*gorm.DB, error) {
	//var dialector = gorm_oracle.Open(dsn)
	var oracleDialector = dialector.(*gorm_oracle.Dialector)
	db, err := gorm.Open(dialector)
	if err != nil {
		return nil, err
	}
	err = db.Callback().Query().Before("gorm:query").
		Register("oracle:check_if_broken_pipe",
			func(tx *gorm.DB) {
				const VersionSQL = "select * from v$version"
				var rawSQL = tx.Statement.SQL.String()
				if rawSQL == VersionSQL {
					return //prevent call loop
				}
				var bannerRecord = new(Banner)
				if err = tx.Debug().Raw(VersionSQL).First(bannerRecord).Error; err == nil {
					return
				}
				fmt.Println(err)
				if !strings.Contains(err.Error(), "broken pipe") {
					return
				}
				fmt.Println("broken pipe detected")
				fmt.Println(err)
				newConnPool, err := sql.Open(dialector.Name(), oracleDialector.DSN)
				if err != nil {
					panic(err)
				}
				//ensure new conn is applied
				db.ConnPool = newConnPool
				db.Statement.ConnPool = newConnPool
				tx.ConnPool = newConnPool
				tx.Statement.ConnPool = newConnPool
			})
	if err != nil {
		panic(err)
	}
	return db, nil
}
