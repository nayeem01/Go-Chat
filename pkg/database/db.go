package database

import (
	"log"

	"github.com/nayeem01/Go-Chat/pkg/config"
	"github.com/nayeem01/Go-Chat/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Db *gorm.DB

func ConncetDB() {

	c := config.Get()
	dbURL := c.GetDBConnStr()
	gDB, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		panic(err)
		// os.Exit(2)
	}
	log.Println("Connected Successfully to PSQL Database")
	gDB.Logger = logger.Default.LogMode(logger.Info)

	gDB.AutoMigrate(&models.User{})
	Db = gDB

}
