package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID              primitive.ObjectID `json:"id" bson:"_id"`
	First_Name      *string            `json:"first_name" validation:"required, min=2, max=30"`
	Last_Name       *string            `json:"last_name" validation:"required, min=2, max=30"`
	Password        *string            `json:"password" validation:"required, min=6"`
	Email           *string            `json:"email" validation:"required, email"`
	Phone           *string            `json:"phone" validation:"required, numeric"`
	Token           *string            `json:"token"`
	Refesh_Token    *string            `json:"refesh_token"`
	Created_At      time.Time          `json:"created_at"`
	Updated_At      time.Time          `json:"updated_at"`
	User_ID         string             `json:"user_id"`
	UserCart        []ProductUser      `json:"usercart" bson:"usercart"`
	Address_Details []Address          `json:"address" bson:"address"`
	Order_Status    []Order            `json:"orders" bson:"orders"`
}

type Product struct {
	Product_ID   primitive.ObjectID `json:"id" bson:"_id"`
	Product_Name *string            `json:"product_name"`
	Price        *uint64            `json:"price"`
	Rating       *uint8             `json:"rating"`
	Image        *string            `json:"image"`
}

type ProductUser struct {
	Product_ID   primitive.ObjectID `bson:"_id"`
	Product_Name *string            `json:"product_name" bson:"product_name"`
	Price        int                `json:"price" bson:"price"`
	Rating       *uint              `json:"rating" bson:"rating"`
	Image        *string            `json:"image" bson:"image"`
}

type Address struct {
	Address_ID primitive.ObjectID `bson:"_id"`
	House_No   *string            `json:"house_no" bson:"house_no"`
	Street     *string            `json:"street" bson:"street"`
	City       *string            `json:"city" bson:"city"`
	Pincode    *string            `json:"pincode" bson:"pincode"`
}

type Order struct {
	Order_ID       primitive.ObjectID `bson:"_id"`
	Order_Cart     []ProductUser      `json:"order_list" bson:"order_list"`
	Ordered_At     time.Time          `json:"ordered_at" bson:"ordered_at"`
	Price          int                `json:"total_price" bson:"total_price"`
	Discount       *int               `json:"discount" bson:"discount"`
	Payment_Method Payment            `json:"payment_method" bson:"payment_method"`
}

type Payment struct {
	Digital bool
	COD     bool
}
