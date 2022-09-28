package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func init() {
	newMigration := &gormigrate.Migration{
		ID: "20220708092000",
		Migrate: func(tx *gorm.DB) error {
			type Comment struct {
				PhotoID *uint64 `valid:"-"`
				AlbumID *uint64 `valid:"-"`
			}
			return tx.AutoMigrate(&Comment{})
		},
		Rollback: func(tx *gorm.DB) error {
			type Comment struct {
				PhotoID uint64 `valid:"-"`
				AlbumID uint64 `valid:"-"`
			}
			return tx.AutoMigrate(&Comment{})
		},
	}

	MigrationsToExec = append(MigrationsToExec, newMigration)
}
