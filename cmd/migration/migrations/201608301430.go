package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
	"time"
	"wallet-api/internal/model"
)

func init() {
	newMigration := &gormigrate.Migration{
		ID: "201608301430",
		Migrate: func(tx *gorm.DB) error {
			type Photo struct {
				ID           uint64      `gorm:"primaryKey;column:id;autoIncrement" valid:"notnull"`
				UrlImage     *string     `valid:"-"`
				AlbumID      uint64      `valid:"notnull"`
				album        model.Album `gorm:"foreignKey:AlbumID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" valid:"notnull"`
				CreatedAt    time.Time   `valid:"-" gorm:"autoCreateTime"`
				DeletedAt    time.Time   `valid:"-"`
				LikeCount    uint64      `valid:"-"`
				CommentCount uint64      `valid:"-"`
			}
			return tx.AutoMigrate(&Photo{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("photos")
		},
	}

	MigrationsToExec = append(MigrationsToExec, newMigration)
}
