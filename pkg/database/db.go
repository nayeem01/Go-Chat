package database

import (
	"log"

	"github.com/nayeem01/Go-Chat/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

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

	Database = DbInstance{
		Db: gDB,
	}
	//db.AutoMigrate(&models.User{})

}
