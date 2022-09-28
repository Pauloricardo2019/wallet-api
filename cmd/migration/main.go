package main

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"log"
	"wallet-api/adapter/database/sql"
	"wallet-api/cmd/migration/migrations"
)

func main() {
	db, err := sql.GetGormDB()
	if err != nil {
		log.Fatal(err)
	}

	migrationsToExec := migrations.GetMigrationsToExec()
	m := gormigrate.New(db, gormigrate.DefaultOptions, migrationsToExec)

	if err = m.Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}
	log.Printf("Migration did run successfully")

}
