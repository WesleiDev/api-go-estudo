package config

import (
	"os"

	"github.com/WesleiDev/api-go-estudo/schemas"
	"gorm.io/driver/sqlite"

	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error){
	logger := GetLogger("sqlite")
	logger.Info("Initializing SQLITE")

	dbPath := "./db/main.db"

	//Check if database file exist

	_, err := os.Stat(dbPath)

	if(os.IsNotExist(err)){
		logger.Info("Database file not found. Creating...")
		//Create the database file and directory
		err = os.MkdirAll("./db", os.ModePerm)
		if(err != nil){
			return nil, err
		}

		file, err := os.Create(dbPath)

		if(err != nil){
			return nil, err
		}
		
		file.Close()
	}

	db, err:= gorm.Open(sqlite.Open(dbPath), &gorm.Config{})

	if(err != nil){
		logger.ErrorF("sqlite opening error %v", err)
		return nil, err
	}

	logger.Info("Running Auto Migration")
	err = db.AutoMigrate(&schemas.Opening{})

	if(err != nil){
		logger.ErrorF("sqlite automigration error: %v", err)
		return nil, err
	}
	logger.Info("Migration successful")

	logger.Info("sqlite opened")

	return db, nil
}