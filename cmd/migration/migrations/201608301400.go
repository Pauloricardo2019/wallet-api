package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
	"time"
)

func init() {
	newMigration := &gormigrate.Migration{
		ID: "201608301400",
		Migrate: func(tx *gorm.DB) error {
			type User struct {
				ID             uint64    `gorm:"primaryKey"`
				UrlImage       *string   `valid:"-"`
				FirstName      string    `valid:"notnull" gorm:"varchar(100)"`
				LastName       string    `valid:"notnull" gorm:"varchar(100)"`
				Email          string    `valid:"notnull" gorm:"varchar(255)"`
				DDD            string    `valid:"-" gorm:"varchar(2)"`
				Phone          string    `valid:"-" gorm:"varchar(9)"`
				Username       string    `valid:"notnull" gorm:"varchar(100)"`
				Birth          time.Time `valid:"-"`
				Biography      string    `valid:"-" gorm:"varchar(255)"`
				Password       string    `valid:"notnull" gorm:"-"`
				HashedPassword string    `valid:"notnull" gorm:"varchar(64)"`
				CreatedAt      time.Time `valid:"-" gorm:"autoCreateTime"`
				UpdatedAt      time.Time `valid:"-" gorm:"autoUpdateTime:milli"`
			}
			return tx.AutoMigrate(&User{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("users")
		},
	}

	MigrationsToExec = append(MigrationsToExec, newMigration)
}
