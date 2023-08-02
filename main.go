package main

import (
	"e-commerce/controllers"
	"e-commerce/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/middleware"
	"google.golang.org/genproto/googleapis/spanner/admin/database/v1"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	app := controllers.NewApplication(database.ProductData(database.Client, "Products"), 
		database.UserData(database.Clint, "Users"))

	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	router.GET("/addtocart", app.AddToCart())
	router.GET("/removeitem", app.RemoveItem())
	router.GET("cartcheckout", app.BuyFromCart())
	router.GET("/instantbuy", app.InstantBuy())

	log.Fatal(router.Run(":" + port))
}
