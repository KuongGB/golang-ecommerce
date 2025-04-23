package routes

import (
	"github.com/gin-gonic/gin"
	"shop/controllers"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/users/signup", controllers.Signup())
	incomingRoutes.POST("/users/login", controllers.Login())
	incomingRoutes.POST("/admin/addproduct", controllers.ProductViewerAdmin())
	incomingRoutes.GET("/users/getproducts", controllers.SearchProduct())
	incomingRoutes.GET("/users/search", controllers.SearchProductByQuery())
	incomingRoutes.POST("/users/addaddress", controllers.AddAddress())
	incomingRoutes.PUT("/users/edithomeaddress", controllers.EditHomeAddress())
	incomingRoutes.PUT("/users/editworkaddress", controllers.EditWorkAddress())
	incomingRoutes.GET("/users/deleteaddress", controllers.DeleteAddress())
}
