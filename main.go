package main

import (
	"fmt"
	AuthController "my-go-wallet/controller/auth"
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

	authorized := r.Group("/users", middleware.JWTAuthen())
	authorized.GET("/getAll", UserController.GetAllUsers)
	authorized.GET("/profile", UserController.Profile)

	authorized = r.Group("/transactions", middleware.JWTAuthen())
	authorized.GET("/getAll", TransactionController.GetAllTransaction)

	r.Use(cors.Default())
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
