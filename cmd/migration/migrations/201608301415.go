package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
	"time"
	"wallet-api/internal/model"
)

func init() {
	newMigration := &gormigrate.Migration{
		ID: "201608301415",
		Migrate: func(tx *gorm.DB) error {
			type Album struct {
				ID           uint64     `valid:"notnull" gorm:"primaryKey; autoIncrement"`
				AlbumCover   *string    `valid:"-"`
				UserID       uint64     `valid:"-"`
				User         model.User `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
				Name         string     `valid:"notnull" gorm:"varchar(100)"`
				Description  string     `valid:"-" gorm:"varchar(255)"`
				LikeCount    uint64     `valid:"-"`
				CommentCount uint64     `valid:"-"`
				CreatedAt    time.Time  `valid:"-" gorm:"autoCreateTime"`
				UpdatedAt    time.Time  `valid:"-" gorm:"autoUpdateTime:milli"`
			}
			return tx.AutoMigrate(&Album{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("albums")
		},
	}

	MigrationsToExec = append(MigrationsToExec, newMigration)
}
