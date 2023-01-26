package backend

import "gorm.io/gorm"

type DBRegistry struct {
	idDBMap map[string]*gorm.DB
}

func NewDBRegistry() *DBRegistry {
	return &DBRegistry{idDBMap: make(map[string]*gorm.DB)}
}

func (x *DBRegistry) Register(id string, db *gorm.DB) {
	x.idDBMap[id] = db
}

func (x *DBRegistry) Get(id string) *gorm.DB {
	db, ok := x.idDBMap[id]
	if !ok {
		return nil
	}
	return db
}
