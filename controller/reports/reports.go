package reports

import (
	"fmt"
	"my-go-wallet/interfaceModel"
	"my-go-wallet/model"
	"my-go-wallet/orm"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetAllTransactionByDate(c *gin.Context) {
	userId := c.MustGet("userId").(float64)

	var transactionInput interfaceModel.TransactionByDateInterface
	if err := c.ShouldBindJSON(&transactionInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var transactions []interfaceModel.TransactionInterface

	fmt.Println(transactionInput.StartDate)
	fmt.Println(transactionInput.EndDate)

	// convert milliseconds to seconds and nanoseconds
	timestampMsStart := int64(transactionInput.StartDate)
	startDate_seconds := timestampMsStart / 1000
	startDate_nanoseconds := (timestampMsStart % 1000) * 1000000
	resultStartDate := time.Unix(startDate_seconds, startDate_nanoseconds)
	dateStart := resultStartDate.Format("2006-01-02 15:04:05")

	timestampMsEnd := int64(transactionInput.EndDate)
	endDate_seconds := timestampMsEnd / 1000
	endDate_nanoseconds := (timestampMsEnd % 1000) * 1000000
	resultEndDate := time.Unix(endDate_seconds, endDate_nanoseconds)
	dateEnd := resultEndDate.Format("2006-01-02 15:04:05")

	fmt.Println(dateStart)
	fmt.Println(dateEnd)

	orm.Db.Model(model.Transaction{}).Where("user_id = ? AND transaction_date BETWEEN ? AND ?", userId, dateStart, dateEnd).Find(&transactions)
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": transactions,
	})
}
