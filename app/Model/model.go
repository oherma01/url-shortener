package model

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

type ShortURL struct {
	ID       uint64 `json:"id" gorm:"primary_key"`
	Redirect string `json:"redirect" gorm:"notnull"`
	Short    string `json:"short-url" gorm:"unique;notnull"`
	Clicked  uint64 `json:"clicked"`
	Random   bool   `json:"random"`

	// TODO: add additional fields
}

func Setup() {
	// use dsn with secrets
	dsn := "" + "user=" + "postgres" + " password=" + "postgres" + " dbname=" + "short-urls" + " port=" + "5432" + " sslmode=disable" + ""

	// raise error if dsn is empty, otherwise connect to database
	var err error

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&ShortURL{})

	if err != nil {
		fmt.Println(err)
	}

}
