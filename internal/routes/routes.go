package routes

import (
	"mamajulia/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	users := r.Group("/auth")
	{
		users.POST("/signin", controllers.Signin)
		users.POST("/signup", controllers.Signup)
	}

	pratos := r.Group("/pratos")
	{
		pratos.GET("/show", controllers.GetPratos)
		pratos.POST("/create", controllers.CreatePratos)
	}
}
