package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// This code is required for other packages to interact with the database. We connect to the database using func Connect
var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "root:silver.shadow!23S@tcp(localhost:3306)/test?charset=utf8&parseTime=True&loc=Local") // OPen connection to the database, specify mysql that is what we need
	// Here, the db we connect to is msql with user askhil and password and the location simplerest. Rest is to identify char and time etc
	if err != nil {
		panic(err)
	}
	db = d // FOr every iteration, that is why we need to save the db data into another var
}

func GetDB() *gorm.DB {
	return db // Whenever this function is invoked, db value is given
}
