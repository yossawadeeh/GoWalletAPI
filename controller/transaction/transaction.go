package transaction

import (
	"my-go-wallet/model"
	"my-go-wallet/orm"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllTransaction(c *gin.Context) {
	userId := c.MustGet("userId").(float64)
	var transaction []model.Transaction

	orm.Db.Find(&transaction, userId)
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": transaction,
	})
}
