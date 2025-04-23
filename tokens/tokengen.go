package tokens

import (
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"log"
	"os"
	"shop/database"
	"time"
)

type SignedDetails struct {
	Email      string `json:"email" bson:"email"`
	First_Name string `json:"first_name" bson:"first_name"`
	Last_Name  string `json:"last_name" bson:"last_name"`
	Uid        string `json:"uid" bson:"uid"`
	jwt.StandardClaims
}

var UserData *mongo.Collection = database.UserData(database.Client, "Users")
var SECRET_KEY = os.Getenv("SECRET_KEY")

func TokenGenerator(email string, first_name string, last_name string, uid string) (signedtoken string, refreshtoken string, err error) {
	claims := &SignedDetails{
		Email:      email,
		First_Name: first_name,
		Last_Name:  last_name,
		Uid:        uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}
	refreshToken := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		return "", "", err
	}
	refreshtoken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, refreshToken).SignedString([]byte(SECRET_KEY))
	if err != nil {
		log.Panic(err)
		return
	}
	return token, refreshtoken, err
}

func ValidateToken(signedtoken string) (claims *SignedDetails, msg string) {
	token, err := jwt.ParseWithClaims(signedtoken, &SignedDetails{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRET_KEY), nil
	})
	if err != nil {
		msg = err.Error()
		return nil, msg
	}
	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		msg = "Token parse error"
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		msg = "Token expired"
		return
	}
	return claims, msg
}

func UpdateAllTokens(signedtoken string, signedrefreshtoken string, userid string) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var updateobj primitive.D
	updateobj = append(updateobj, bson.E{Key: "token", Value: signedtoken})
	updateobj = append(updateobj, bson.E{Key: "refresh_token", Value: signedrefreshtoken})
	update_at, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	updateobj = append(updateobj, bson.E{Key: "update_at", Value: update_at})
	upsert := true
	filter := bson.M{"user_id": userid}
	opt := options.UpdateOptions{
		Upsert: &upsert,
	}
	_, err := UserData.UpdateOne(ctx, filter, bson.D{
		{Key: "$set", Value: updateobj},
	}, &opt)
	defer cancel()
	if err != nil {
		log.Panic(err)
		return
	}
}
