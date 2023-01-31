package v1

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	// group for auth functions

	auth := r.Group("/api/v1/auth")
	{

		// common auth functions

		auth.POST("/login", h.Login)
		auth.POST("/registration", h.Registration)

		// token from email functions

		auth.POST("/forgotPassword", h.ForgotPassword)
		auth.POST("/resetPassword/:token", h.ResetPassword)
		auth.GET("/verifyEmail/:token", h.Verification)

		// already authorized functions

		auth.GET("/logout", h.CheckTokenAuthHandler(h.RequestHandler(h.Logout)))       // delete Authentication session
		auth.GET("/token", h.CheckTokenAuthHandler(h.RequestHandler(GetTokenByToken))) // set new Authorization Token from Authentication Token
	}

	common := r.Group("/api/v1/user")
	{
		common.GET("/getSessions", h.CheckTokenHandler(h.GetSessions))
		common.GET("/getBins", h.CheckTokenHandler(h.GetBins))
		common.POST("/deleteSessions", h.CheckTokenHandler(h.DeleteSessions))
		common.POST("/changeUserInfo", h.CheckTokenHandler(h.ChangeUserInfo))
		common.POST("/repeatTask", h.CheckTokenHandler(h.RepeatTask))
		common.POST("/cancelTask", h.CheckTokenHandler(h.RepeatTask))
		common.POST("/createTaskCardCreation", h.CheckTokenHandler(h.CreateTaskCardCreation))
		common.POST("/createTaskCardReplenishment", h.CheckTokenHandler(h.CreateTaskCardReplenishment))
		common.GET("/getCards", h.CheckTokenHandler(h.GetCards))                     // get user cards from db
		common.GET("/getTasks", h.CheckTokenHandler(h.GetTasks))                     // get user tasks from db
		common.GET("/getUserInfo", h.CheckTokenHandler(h.GetUserInfo))               // get user info
		common.GET("/getTransactions", h.CheckTokenHandler(h.GetTransactions))       // get transactions from card provider
		common.GET("/getAllTransactions", h.CheckTokenHandler(h.GetAllTransactions)) // get all transactions
		common.GET("/getSuppliers", h.CheckTokenHandler(h.GetSuppliers))             // get all transactions
	}

	// group for admin

	admin := r.Group("api/v1/admin")
	{

		// common auth functions

		admin.POST("/login", h.AdminLogin)

		// already authorized functions

		admin.GET("/logout", h.AdminCheckTokenAuthHandler(h.RequestHandler(h.Logout)))            // delete Authentication session
		admin.GET("/token", h.AdminCheckTokenAuthHandler(h.RequestHandler(AdminGetTokenByToken))) // set new Authorization Token from Authentication Token

		// get access to another account functions

		admin.POST("/getAccess", h.AdminCheckTokenHandler(h.GetAccess)) // set token, redirect from frontend
		admin.GET("/access", h.CheckTokenAuthHandler(h.Access))

		// get info

		admin.GET("/getSessions", h.AdminCheckTokenHandler(h.GetSessions))
		admin.GET("/getUsers", h.AdminCheckTokenHandler(h.AdminGetUsers))                     // get all users from db
		admin.GET("/getCards", h.AdminCheckTokenHandler(h.AdminGetCards))                     // get all cards from db
		admin.GET("/getTasks", h.AdminCheckTokenHandler(h.AdminGetTasks))                     // get all tasks from db
		admin.GET("/getSuppliers", h.AdminCheckTokenHandler(h.GetSuppliers))                  // get all connected suppliers
		admin.GET("/getAllTransactions", h.AdminCheckTokenHandler(h.AdminGetAllTransactions)) // get all transactions
		admin.GET("/getFinLogs", h.AdminCheckTokenHandler(h.AdminGetFinLogs))                 // get logs
		admin.POST("/getActionLog", h.AdminCheckTokenHandler(h.AdminGetActionLog))            // get action log from main log id
		admin.POST("/getTransactions", h.AdminCheckTokenHandler(h.GetTransactions))           // get transactions from card provider

		// edit info

		admin.POST("/addUser", h.AdminCheckTokenHandler(h.AdminAddUser))                             // add user
		admin.POST("/deleteUser", h.AdminCheckTokenHandler(h.DeleteUser))                            // delete user
		admin.POST("/updateUser", h.AdminCheckTokenHandler(h.AdminUpdateUser))                       // update user info
		admin.POST("/updateUsersCommission", h.AdminCheckTokenHandler(h.AdminUpdateUsersCommission)) // update user info
		admin.POST("/topUpBalance", h.AdminCheckTokenHandler(h.AdminTopUpBalance))                   // Top up user balance
		admin.POST("/topUpCardBalance", h.AdminCheckTokenHandler(h.TopUpCardBalance))                // Top up card balance
		admin.POST("/changeCardStatus", h.AdminCheckTokenHandler(h.ChangeCardStatus))                // change card status
		admin.POST("/changeCardInfo", h.AdminCheckTokenHandler(h.ChangeCardInfo))                    // change card info
		admin.POST("/changeCardLimit", h.AdminCheckTokenHandler(h.ChangeCardLimit))                  // change card limit
		admin.POST("/createCards", h.AdminCheckTokenHandler(h.CreateCards))                          // create cards
		admin.POST("/createCardByTask", h.AdminCheckTokenHandler(h.CardCreationByTask))              // create card by task
		admin.POST("/cardReplenishmentByTask", h.AdminCheckTokenHandler(h.CardReplenishmentByTask))  // Top up card balance by task
		admin.POST("/updateSupplier", h.AdminCheckTokenHandler(h.UpdateSupplier))                    // update supplier info

		admin.POST("/deleteSessions", h.AdminCheckTokenHandler(h.DeleteSessions))
	}
}
