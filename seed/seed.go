package main

import (
	"fmt"
	"my-go-wallet/model"
	"my-go-wallet/orm"

	"github.com/joho/godotenv"
)

func main() {

	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	orm.InitDB()

	// Migrate the schema
	orm.Db.AutoMigrate(&model.User{}, &model.TypeTransaction{}, &model.CategoryTransaction{}, &model.Transaction{})

	// Create
	typeIncome := model.TypeTransaction{TypeName: "Income"}
	orm.Db.Create(&typeIncome)
	typeExpenses := model.TypeTransaction{TypeName: "Expenses"}
	orm.Db.Create(&typeExpenses)

	orm.Db.Model(&model.CategoryTransaction{}).Create([]map[string]interface{}{
		{"CategoryName": "Food", "TypeID": typeExpenses.ID},
		{"CategoryName": "Shopping", "TypeID": typeExpenses.ID},
		{"CategoryName": "Bill", "TypeID": typeExpenses.ID},
		{"CategoryName": "Travel", "TypeID": typeExpenses.ID},
		{"CategoryName": "Transportation", "TypeID": typeExpenses.ID},
		{"CategoryName": "Investment", "TypeID": typeExpenses.ID},
		{"CategoryName": "OtherExpenses", "TypeID": typeExpenses.ID},

		{"CategoryName": "Salary", "TypeID": typeIncome.ID},
		{"CategoryName": "Bonus", "TypeID": typeIncome.ID},
		{"CategoryName": "Reward", "TypeID": typeIncome.ID},
		{"CategoryName": "OtherIncome", "TypeID": typeIncome.ID},
	})

}
