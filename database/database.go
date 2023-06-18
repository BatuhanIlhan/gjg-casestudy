package database

import (
	"database/sql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"log"
)

type Config struct {
	URI             string
	MigrationSource string
	Debug           bool
}

type Database struct {
	db     *sql.DB
	config Config
}

func New(config Config) *Database {
	return &Database{config: config}
}

func (d *Database) Connect() {
	var err error
	if d.db, err = sql.Open("postgres", d.config.URI); err != nil {
		log.Fatalf("[Database] Connection failed: %v", err.Error())
	}
	boil.DebugMode = d.config.Debug
}

func (d *Database) Close() {
	if d.db == nil {
		return
	}
	if err := d.db.Close(); err != nil {
		log.Printf("[Database] Close failed: %v", err.Error())
	}
}

func (d *Database) DB() *sql.DB {
	return d.db
}

func (d *Database) Migrate() {
	if driver, err := postgres.WithInstance(d.db, &postgres.Config{}); err != nil {
		log.Fatalf("[Database] Migration driver could not initialized: %v", err.Error())
	} else if m, err := migrate.NewWithDatabaseInstance(d.config.MigrationSource, "", driver); err != nil {
		log.Fatalf("[Database] Migration instance could not initialized: %v", err.Error())
	} else if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("[Database] Migration failed: %v", err.Error())
	}
}
