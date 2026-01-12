package routes

import (
	"mamajulia/internal/controllers"
	"mamajulia/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	auth := r.Group("/auth")
	{
		auth.POST("/signup", controllers.Signup)
		auth.POST("/signin", controllers.Signin)
	}

	pratos := r.Group("/pratos")
	{
		pratos.GET("", controllers.GetPratos)
		pratos.GET("/:id", controllers.GetPratoByID)
	}

	protected := r.Group("/")
	protected.Use(middlewares.AuthMiddleware())
	{
		pedidos := protected.Group("/pedidos")
		{
			pedidos.POST("", controllers.CreatePedido)
		}

		admin := protected.Group("/admin")
		admin.Use(middlewares.AdminMiddleware())
		{
			admin.POST("/pratos", controllers.CreatePratos)
			admin.PUT("/pratos/:id", controllers.UpdatePratos)
			admin.DELETE("/pratos/:id", controllers.DeletePratos)

			admin.GET("/pedidos", controllers.GetPedidos)
			admin.GET("/pedidos/:id", controllers.GetPedidoByID)
			admin.PUT("/pedidos/:id/status", controllers.UpdatePedidoStatus)
			admin.DELETE("/pedidos/:id", controllers.DeletePedido)
		}
	}
}
