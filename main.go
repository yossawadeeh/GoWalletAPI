package main

import (
	"fmt"
	AuthController "my-go-wallet/controller/auth"
	ReportsController "my-go-wallet/controller/reports"
	TransactionController "my-go-wallet/controller/transaction"
	UserController "my-go-wallet/controller/user"
	"my-go-wallet/middleware"
	"my-go-wallet/orm"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	orm.InitDB()
	r := gin.Default()

	r.POST("/register", AuthController.Register)
	r.POST("/login", AuthController.Login)

	userGroup := r.Group("/users", middleware.JWTAuthen())
	userGroup.GET("/getAllUsers", UserController.GetAllUsers)
	userGroup.GET("/getUser", UserController.Profile)

	transactionsGroup := r.Group("/transactions", middleware.JWTAuthen())
	transactionsGroup.GET("/getAllTransactions", TransactionController.GetAllTransactions)
	transactionsGroup.POST("/insertTransaction", TransactionController.CreateTransaction)
	transactionsGroup.DELETE("/deleteTransaction/:id", TransactionController.DeleteTransaction)
	transactionsGroup.PUT("/updateTransaction/:id", TransactionController.UpdateTransaction)

	reportGroup := r.Group("/reports", middleware.JWTAuthen())
	reportGroup.POST("/getAllTransactionsByDate", ReportsController.GetAllTransactionByDate)

	r.Use(cors.Default())
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
