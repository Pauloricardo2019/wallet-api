package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
	"time"
	"wallet-api/internal/model"
)

func init() {
	newMigration := &gormigrate.Migration{
		ID: "201608301490",
		Migrate: func(tx *gorm.DB) error {
			type DeletedPhoto struct {
				ID        uint64      `gorm:"primaryKey;column:id;autoIncrement" valid:"notnull"`
				PhotoID   uint64      `valid:"notnull"`
				photo     model.Photo `gorm:"foreignKey:PhotoID;references:ID" valid:"notnull"`
				AlbumID   uint64      `valid:"notnull"`
				album     model.Album `gorm:"foreignKey:AlbumID;references:ID" valid:"notnull"`
				UserID    uint64      `valid:"notnull "`
				user      model.User  `gorm:"foreignKey:UserID;references:ID" valid:"notnull"`
				DeletedAt time.Time   `valid:"notnull "`
			}
			return tx.AutoMigrate(&DeletedPhoto{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("deleted_photos")
		},
	}

	MigrationsToExec = append(MigrationsToExec, newMigration)
}
