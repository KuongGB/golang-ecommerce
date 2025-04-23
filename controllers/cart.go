package controllers

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"shop/database"
	"shop/models"
	"time"
)

type Application struct {
	prodCollection *mongo.Collection
	userCollection *mongo.Collection
}

func NewApplication(prodCollection, userCollection *mongo.Collection) *Application {
	return &Application{
		prodCollection: prodCollection,
		userCollection: userCollection,
	}
}

func (app *Application) AddToCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		productQueryID := c.Query("id")
		if productQueryID == "" {
			log.Println("No product ID provided")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("No product ID provided"))
			return
		}
		userQueryID := c.Query("userID")
		if userQueryID == "" {
			log.Println("No user ID provided")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("No user ID provided"))
			return
		}
		productID, err := primitive.ObjectIDFromHex(productQueryID)
		if err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err = database.AddProductToCart(ctx, app.prodCollection, app.userCollection, productID, userQueryID)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
		}
		c.IndentedJSON(http.StatusOK, "Product added to cart")
	}
}

func (app *Application) RemoteItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		productQueryID := c.Query("id")
		if productQueryID == "" {
			log.Println("No product ID provided")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("No product ID provided"))
			return
		}
		userQueryID := c.Query("userID")
		if userQueryID == "" {
			log.Println("No user ID provided")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("No user ID provided"))
			return
		}
		productID, err := primitive.ObjectIDFromHex(productQueryID)
		if err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err = database.RemoveCartItem(ctx, app.prodCollection, app.userCollection, productID, userQueryID)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
			return
		}
		c.IndentedJSON(http.StatusOK, "Product removed from cart")
	}

}

func (app *Application) GetItemFromCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		user_id := c.Query("id")
		if user_id == "" {
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusNotFound, gin.H{"error": "No user ID provided"})
			c.Abort()
			return
		}
		usert_id, _ := primitive.ObjectIDFromHex(user_id)
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var filtedcart models.User
		err := UserCollection.FindOne(ctx, bson.D{primitive.E{Key: "_id", Value: usert_id}}).Decode(&filtedcart)

		if err != nil {
			log.Println(err)
			c.IndentedJSON(500, "No user found")
			return
		}
		filter_match := bson.D{{Key: "$match", Value: bson.D{primitive.E{Key: "_id", Value: usert_id}}}}
		unwind := bson.D{{Key: "$unwind", Value: bson.D{primitive.E{Key: "path", Value: "$usercart"}}}}
		grouping := bson.D{{Key: "$group", Value: bson.D{primitive.E{Key: "_id", Value: "$_id"}, {Key: "total", Value: bson.D{primitive.E{Key: "$sum", Value: "$usercart.price"}}}}}}
		pointcursor, err := UserCollection.Aggregate(ctx, mongo.Pipeline{filter_match, unwind, grouping})
		if err != nil {
			log.Println(err)
		}
		var listing []bson.M
		if err = pointcursor.All(ctx, &listing); err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
		}
		for _, json := range listing {
			c.IndentedJSON(200, json["total"])
			c.IndentedJSON(http.StatusOK, filtedcart.UserCart)
		}
		ctx.Done()
	}
}

func (app *Application) BuyFromCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		userQueryID := c.Query("id")
		if userQueryID == "" {
			log.Panic("No user ID provided")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("No user ID provided"))
			return
		}
		var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		err := database.BuyItemFromCart(ctx, app.userCollection, userQueryID)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
		}

		c.IndentedJSON(http.StatusOK, "Product buy from cart")
	}
}

func (app *Application) InstantBuy() gin.HandlerFunc {
	return func(c *gin.Context) {
		productQueryID := c.Query("id")
		if productQueryID == "" {
			log.Println("No product ID provided")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("No product ID provided"))
			return
		}
		userQueryID := c.Query("userID")
		if userQueryID == "" {
			log.Println("No user ID provided")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("No user ID provided"))
			return
		}
		productID, err := primitive.ObjectIDFromHex(productQueryID)
		if err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err = database.InstantBuyer(ctx, app.prodCollection, app.userCollection, productID, userQueryID)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
			return
		}
		c.IndentedJSON(http.StatusOK, "Ordered successfully")
	}
}
