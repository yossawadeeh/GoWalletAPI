package auth

import (
	"fmt"
	"my-go-wallet/model"
	"my-go-wallet/orm"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var hmacSampleSecret []byte

func Register(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	encryptPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	register := model.User{Username: user.Username, Password: string(encryptPassword), Fullname: user.Fullname, Email: user.Email, Avatar: user.Avatar}

	var userExist model.User
	orm.Db.Where("username = ?", register.Username).First(&userExist)
	if userExist.ID > 0 {
		c.JSON(http.StatusOK, gin.H{
			"status":  "error",
			"message": "User is exist",
		})
		return
	}

	result := orm.Db.Create(&register)
	c.JSON(http.StatusOK, gin.H{
		"status":      "ok",
		"RowAffected": result.RowsAffected,
	})
}

func Login(c *gin.Context) {
	var user model.Login
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	encryptPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	login := model.Login{Username: user.Username, Password: string(encryptPassword)}

	var userExist model.User
	orm.Db.Where("username = ?", login.Username).First(&userExist)
	if userExist.ID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"status":  "error",
			"message": "User does not exist",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(userExist.Password), []byte(user.Password))
	if err == nil {
		hmacSampleSecret = []byte(os.Getenv("JWT_SECRET_KEY"))
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"userId": userExist.ID,
			"exp":    time.Now().Add(time.Hour * 1).Unix(),
		})
		// Sign and get the complete encoded token as a string using the secret
		tokenString, err := token.SignedString(hmacSampleSecret)
		fmt.Println(tokenString, err)

		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "Login success",
			"token":   tokenString,
		})

	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  "error",
			"message": "Login failed",
		})
	}

}
