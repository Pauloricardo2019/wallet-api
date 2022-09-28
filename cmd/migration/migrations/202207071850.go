package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func init() {
	newMigration := &gormigrate.Migration{
		ID: "20220707185000",
		Migrate: func(tx *gorm.DB) error {
			type Like struct {
				PhotoID *uint64 `valid:"-"`
				AlbumID *uint64 `valid:"-"`
			}
			return tx.AutoMigrate(&Like{})
		},
		Rollback: func(tx *gorm.DB) error {
			type Like struct {
				PhotoID uint64 `valid:"-"`
				AlbumID uint64 `valid:"-"`
			}
			return tx.AutoMigrate(&Like{})
		},
	}

	MigrationsToExec = append(MigrationsToExec, newMigration)
}
