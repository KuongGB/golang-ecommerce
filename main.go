package main

import (
	_ "fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"log"
	_ "net/http"
	"os"
	"shop/controllers"
	"shop/database"
	_ "shop/database"
	"shop/middleware"
	_ "shop/middleware"
	"shop/routes"
	_ "shop/routes"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	app := controllers.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))

	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	router.GET("/addtocart", app.AddToCart())
	router.GET("/removecartitem", app.RemoteItem())
	router.GET("/cartcheckout", app.BuyFromCart())
	router.GET("/instantbuy", app.InstantBuy())
	router.GET("/listcart", app.GetItemFromCart())

	log.Fatal(router.Run(":" + port))
}
