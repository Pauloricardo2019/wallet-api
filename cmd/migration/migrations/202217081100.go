package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
	"time"
	"wallet-api/internal/model"
)

func init() {
	newMigration := &gormigrate.Migration{
		ID: "20221708110000",
		Migrate: func(tx *gorm.DB) error {
			type Feed struct {
				ID        uint64      `gorm:"primaryKey;autoIncrement"`
				UserID    uint64      `valid:"notnull"`
				User      model.User  `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
				Action    string      `valid:"notnull"`
				PhotoID   uint64      `valid:"notnull"`
				Photo     model.Photo `gorm:"foreignKey:PhotoID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
				CreatedAt time.Time   `gorm:"autoCreateTime"`
			}
			return tx.AutoMigrate(&Feed{})
		},
		Rollback: func(tx *gorm.DB) error {

			return tx.Migrator().DropTable("feeds")
		},
	}

	MigrationsToExec = append(MigrationsToExec, newMigration)
}
