package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	dbUser string
	dbPswd string
	dbHost string
	dbPort string
	dbName string
}

func Get() *Config {

	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}

	conf := &Config{dbUser: os.Getenv("POSTGRES_USER"), dbPswd: os.Getenv("POSTGRES_PASSWORD"),
		dbHost: os.Getenv("POSTGRES_HOST"), dbPort: os.Getenv("POSTGRES_PORT"), dbName: os.Getenv("POSTGRES_DB")}

	return conf
}

func (c *Config) GetDBConnStr() string {

	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		c.dbUser,
		c.dbPswd,
		c.dbHost,
		c.dbPort,
		c.dbName,
	)
}
