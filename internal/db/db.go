package db

import (

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
    DB *gorm.DB
)

func ConnectDB(dsn string) error {
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        return err
    }

    DB = db
    return nil
}


