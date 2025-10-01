package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID             primitive.ObjectID `json:"_id" bson:"_id"`
	FirstName      *string            `json:"first_name" bson:"first_name" validate:"required,min=2,max=30"`
	LastName       *string            `json:"last_name" bson:"last_name" validate:"required,min=2,max=30"`
	Password       *string            `json:"password" bson:"password" validate:"required,min=6,max=100"`
	Email          *string            `json:"email" bson:"email" validate:"required,email"`
	Phone          *string            `json:"phone" bson:"phone" validate:"required,len=10,numeric"`
	Token          *string            `json:"token" bson:"token"`
	RefreshToken   *string            `json:"refresh_token" bson:"refresh_token"`
	CreatedAt      time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt      time.Time          `json:"updated_at" bson:"updated_at"`
	UserID         *string            `json:"user_id" bson:"user_id"`
	UserCart       []ProductUser      `json:"user_cart" bson:"user_cart"`
	AddressDetails []Address          `json:"address_details" bson:"address_details"`
	OrderStatus    []Order            `json:"order_status" bson:"order_status"`
}

type Product struct {
	ProductID   primitive.ObjectID `bson:"_id"`
	ProductName *string            `json:"product_name" bson:"product_name" validate:"required"`
	Price       *uint64            `json:"price" bson:"price" validate:"required,gte=1"`
	Rating      *uint8             `json:"rating" bson:"rating" validate:"omitempty,gte=0,lte=5"`
	Image       *string            `json:"image" bson:"image" validate:"omitempty,url"`
}

type ProductUser struct {
	ProductID   primitive.ObjectID `bson:"_id"`
	ProductName *string            `json:"product_name" bson:"product_name" validate:"required"`
	Price       int                `json:"price" bson:"price" validate:"required"`
	Rating      *uint              `json:"rating" bson:"rating" validate:"omitempty,gte=0,lte=5"`
	Image       *string            `json:"image" bson:"image" validate:"omitempty,url"`
}

type Address struct {
	AddressID primitive.ObjectID `json:"_id" bson:"_id"`
	House     *string            `json:"house" bson:"house" validate:"required"`
	Street    *string            `json:"street" bson:"street" validate:"required"`
	City      *string            `json:"city" bson:"city" validate:"required"`
	Pincode   *string            `json:"pincode" bson:"pincode" validate:"required,len=6,numeric"`
}

type Order struct {
	OrderID       primitive.ObjectID `json:"_id" bson:"_id"`
	OrderCart     []ProductUser      `json:"order_cart" bson:"order_cart" validate:"required,dive"`
	OrderedAt     time.Time          `json:"ordered_at" bson:"ordered_at"`
	Price         int                `json:"price" bson:"price" validate:"required,gte=1"`
	Discount      *int               `json:"discount" bson:"discount" validate:"omitempty,gte=0,lte=100"`
	PaymentMethod Payment            `json:"payment_method" bson:"payment_method" validate:"required"`
}

type Payment struct {
	Digital bool `json:"digital" bson:"digital"`
	COD     bool `json:"cod" bson:"cod"`
}
