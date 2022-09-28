package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
	"time"
	"wallet-api/internal/model"
)

func init() {
	newMigration := &gormigrate.Migration{
		ID: "201608301460",
		Migrate: func(tx *gorm.DB) error {
			type Comment struct {
				ID        uint64       `gorm:"primaryKey;column:id;autoIncrement" valid:"-"`
				Comment   string       `valid:"-"`
				PhotoID   uint64       `valid:"-"`
				Photo     *model.Photo `gorm:"foreignKey:PhotoID" valid:"-"`
				AlbumID   uint64       `valid:"-"`
				Album     *model.Album `gorm:"foreignKey:AlbumID" valid:"-"`
				UserID    uint64       `valid:"-"`
				CreatedAt time.Time    `valid:"-" gorm:"autoCreateTime"`
			}
			return tx.AutoMigrate(&Comment{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("comments")
		},
	}

	MigrationsToExec = append(MigrationsToExec, newMigration)
}
