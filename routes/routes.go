package routes

import (
	"github.com/ashish00304/test-test/controllers"
	"github.com/ashish00304/test-test/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {

	user := r.Group("/users")
	{
		user.GET("/printHello", controllers.PrintHello)
		user.POST("/addUser", controllers.AddUser)
		user.GET("/getAllUser", controllers.GetAllUser)
		user.DELETE("/deleteUser/:id", controllers.DeleteUser)
		user.PATCH("/updateDetails/:id", controllers.UpdateUserDetails)
		user.POST("/login", controllers.Login)

	}

	secureUser := r.Group("/secure")
	secureUser.Use(middleware.AuthMiddleware())
	{
		secureUser.GET("/profile", controllers.SecureProfile)
	}

}
