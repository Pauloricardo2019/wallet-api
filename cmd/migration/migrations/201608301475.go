package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
	"time"
	"wallet-api/internal/model"
)

func init() {
	newMigration := &gormigrate.Migration{
		ID: "201608301475",
		Migrate: func(tx *gorm.DB) error {
			type Token struct {
				Value     string `valid:"notnull" gorm:"varchar(255)"`
				UserID    uint64
				User      model.User `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" valid:"notnull"`
				CreatedAt time.Time  `valid:"-" gorm:"autoCreateTime"`
				ExpiresAt time.Time  `valid:"-" gorm:"autoUpdateTime:milli"`
			}
			return tx.AutoMigrate(&Token{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("tokens")
		},
	}

	MigrationsToExec = append(MigrationsToExec, newMigration)
}
