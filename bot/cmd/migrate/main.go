package main

import (
	"errors"
	"fmt"
	"github.com/deed-labs/openroll/bot/configs"
	"github.com/deed-labs/openroll/bot/migrations"
	"github.com/golang-migrate/migrate/v4"
	"log"
	"net/url"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	bindata "github.com/golang-migrate/migrate/v4/source/go_bindata"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println(fmt.Errorf("no .env file found"))
	}

	conf, err := configs.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	u, err := url.Parse(conf.PostgresDSN)
	if err != nil {
		log.Fatal(err)
	}
	sqlConnString := u.String()

	if err := migrateDB(sqlConnString); err != nil {
		log.Fatal(err)
	}
}

func migrateDB(sqlConnString string) error {
	s := bindata.Resource(migrations.AssetNames(), migrations.Asset)

	d, err := bindata.WithInstance(s)
	if err != nil {
		return err
	}
	m, err := migrate.NewWithSourceInstance("go-bindata", d, sqlConnString)
	if err != nil {
		return err
	}
	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Printf("[Migrations] Error: [%v]", err)
	}

	version, dirty, err := m.Version()
	if err != nil {
		return err
	}
	if dirty {
		return errors.New("migrations is dirty")
	}
	log.Printf("Migration Version = %d \n", version)
	return nil
}
