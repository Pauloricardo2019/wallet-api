package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
	"time"
	"wallet-api/internal/model"
)

func init() {
	newMigration := &gormigrate.Migration{
		ID: "201608301445",
		Migrate: func(tx *gorm.DB) error {
			type Like struct {
				ID        uint64       `gorm:"primaryKey;column:id;autoIncrement" valid:"-"`
				PhotoID   uint64       `valid:"-"`
				Photo     *model.Photo `gorm:"foreignKey:PhotoID;references:ID;constraint:OnDelete:CASCADE;" valid:"-"`
				AlbumID   uint64       `valid:"-"`
				Album     *model.Album `gorm:"foreignKey:AlbumID;references:ID;constraint:OnDelete:CASCADE;" valid:"-"`
				UserID    uint64       `valid:"-"`
				CreatedAt time.Time    `valid:"-" gorm:"autoCreateTime"`
			}
			return tx.AutoMigrate(&Like{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("likes")
		},
	}

	MigrationsToExec = append(MigrationsToExec, newMigration)
}
