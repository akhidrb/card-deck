package tests

import (
	"fmt"
	"github.com/cucumber/godog"
	"github.com/golang-migrate/migrate/v4"
	postgresMigrate "github.com/golang-migrate/migrate/v4/database/postgres"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

var testConfig TestConfig

type TestConfig struct {
	adapter Client
}

func InitTestSuite(ctx *godog.TestSuiteContext) {
	clientInst := NewClient("http://localhost:8080/api/v1", time.Second*10)
	testConfig = TestConfig{adapter: clientInst}
	dsn := fmt.Sprintf("host=localhost user=toggl password=toggl dbname=cards port=5432 sslmode=disable")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	runMigrations(db)
}

func runMigrations(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	driver, err := postgresMigrate.WithInstance(sqlDB, &postgresMigrate.Config{})
	if err != nil {
		log.Fatal(err)
	}
	m, err := migrate.NewWithDatabaseInstance("file://../migrations", "postgres", driver)
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Down(); err != nil {
		log.Println("Database wasn't dropped", "err", err.Error())
	}
	if err := m.Up(); err != nil {
		log.Println("Migrations weren't updated", "err", err.Error())
	}
}
