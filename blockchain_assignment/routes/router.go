package routes

import (
	"atm-system/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoute(r *gin.Engine) {
	r.POST("/CreateAccount", controllers.CreateAccount)
	r.POST("/DepositMoney", controllers.Deposit)
	r.POST("/Withdrawal", controllers.Withdraw)
	r.POST("/TransferMoney", controllers.Transfer)
	r.POST("/SetPin", controllers.SetPin)
	r.POST("/BankStatement", controllers.BankStatement)
}
