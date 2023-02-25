package main

import (
	"my-go-wallet/model"
	"my-go-wallet/orm"
)

func main() {

	// Migrate the schema
	orm.Db.AutoMigrate(&model.User{})

	// Create
	//db.Create(&model.User{Username: "yeolowbatt", Password: "123"})

}
