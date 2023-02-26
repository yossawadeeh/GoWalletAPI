package transaction

import (
	"my-go-wallet/interfaceModel"
	"my-go-wallet/model"
	"my-go-wallet/orm"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetAllTransactions(c *gin.Context) {
	userId := c.MustGet("userId").(float64)
	var transactions []interfaceModel.TransactionInterface

	orm.Db.Model(model.Transaction{}).Where("user_id = ?", userId).Find(&transactions)
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": transactions,
	})
}

func CreateTransaction(c *gin.Context) {
	userId := c.MustGet("userId").(float64)

	var transaction interfaceModel.TransactionInterface
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newTransaction := model.Transaction{
		TransactionDate: time.Now(),
		Total:           transaction.Total,
		Note:            transaction.Note,
		CategoryID:      transaction.CategoryID,
		UserID:          userId,
	}

	var category model.CategoryTransaction
	orm.Db.Where("id = ?", newTransaction.CategoryID).First(&category)
	if category.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Not found category",
		})
		return
	}

	result := orm.Db.Create(&newTransaction)

	c.JSON(http.StatusOK, gin.H{
		"status":         "ok",
		"RowAffected":    result.RowsAffected,
		"newTransaction": transaction,
	})
}

func DeleteTransaction(c *gin.Context) {
	//userId := c.MustGet("userId").(float64)

	transactionId := c.Param("id")
	var transaction model.Transaction
	orm.Db.First(&transaction, transactionId)
	if transaction.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Not found transaction",
		})
		return
	}

	result := orm.Db.Delete(&transaction)
	c.JSON(http.StatusOK, gin.H{
		"status":        "ok",
		"RowAffected":   result.RowsAffected,
		"TransactionId": transaction.ID,
	})

}

func UpdateTransaction(c *gin.Context) {
	transactionId := c.Param("id")
	var transaction interfaceModel.UpdateTransactionInterface
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var transactionQuery model.Transaction
	orm.Db.First(&transactionQuery, transactionId)
	if transactionQuery.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Not found transaction",
		})
		return
	}

	var category model.CategoryTransaction
	orm.Db.Where("id = ?", transaction.CategoryID).First(&category)
	if category.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Not found category",
		})
		return
	}

	updateTime := transactionQuery.TransactionDate
	var err error
	if transaction.TransactionDate != "" {
		updateTime, err = time.Parse("2006-01-02 15:04:05.000", transaction.TransactionDate)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"message": err.Error(),
			})
			return
		}
	}

	updateNote := transactionQuery.Note
	if transaction.Note != "" {
		updateNote = transaction.Note
	}

	transactionQuery.CategoryID = transaction.CategoryID
	transactionQuery.TransactionDate = updateTime
	transactionQuery.Total = transaction.Total
	transactionQuery.Note = updateNote

	result := orm.Db.Save(transactionQuery)
	c.JSON(http.StatusOK, gin.H{
		"status":        "ok",
		"RowAffected":   result.RowsAffected,
		"TransactionId": transactionQuery.ID,
	})

}
