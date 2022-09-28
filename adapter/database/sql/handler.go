package sql

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"sync"
	"wallet-api/adapter/config"
)

var gormDb *gorm.DB
var mutexDB sync.Mutex

func GetGormDB() (*gorm.DB, error) {
	mutexDB.Lock()
	defer mutexDB.Unlock()

	if gormDb != nil {
		return gormDb, nil
	}

	newDb, err := gorm.Open(postgres.Open(config.GetConfig().DbConnString), &gorm.Config{FullSaveAssociations: true})
	if err != nil {
		return nil, err
	}

	gormDb = newDb

	return gormDb, nil
}
