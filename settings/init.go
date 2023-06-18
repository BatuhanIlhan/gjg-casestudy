package settings

import (
	"github.com/Netflix/go-env"
	"log"
	"time"
)

var Cnf config

type config struct {
	IdleTimeout     time.Duration `env:"IDLE_TIME_OUT"`
	ReadTimeout     time.Duration `env:"READ_TIME_OUT"`
	WriteTimeout    time.Duration `env:"WRITE_TIME_OUT"`
	PostgresURI     string        `env:"POSTGRES_URI"`
	MigrationSource string        `env:"MIGRATION_SOURCE"`
	Debug           bool          `env:"DEBUG"`
	BaseUrl         string        `env:"BASEURL"`
	Port            string        `env:"PORT"`
}

func init() {
	_, err := env.UnmarshalFromEnviron(&Cnf)
	if err != nil {
		log.Fatal(err)
	}
}
