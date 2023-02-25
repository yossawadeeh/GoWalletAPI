package user

import (
	"my-go-wallet/model"
	"my-go-wallet/orm"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	var users []model.User
	orm.Db.Find(&users)
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": users,
	})
}

func Profile(c *gin.Context) {
	// retrieve userId from token and find profile

	userId := c.MustGet("userId").(float64)
	var user model.User
	orm.Db.First(&user, userId)
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": user,
	})
}
