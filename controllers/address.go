package controllers

import (
	"context"
	"ecommerce/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddAddress() gin.HandlerFunc {
	return func(c *gin.Context) {
		userid := c.Query("_id")
		if userid == "" {
			c.Header("content-Type", "application/json")
			c.JSON(http.StatusNotFound, gin.H{"Error": "Invalid serach index"})
			c.Abort()
			return
		}
		usertid, err := primitive.ObjectIDFromHex(userid)
		if err != nil {
			c.IndentedJSON(500, "Internal server error")
		}
		var addresses models.Address
		addresses.AddressID = primitive.NewObjectID()
		if err = c.BindJSON(&addresses); err != nil {
			c.IndentedJSON(http.StatusNotAcceptable, err.Error())
		}
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		matchfilter := bson.D{{Key: "$match", Value: bson.D{primitive.E{Key: "_id", Value: usertid}}}}
		unwind := bson.D{{Key: "$unwind", Value: bson.D{primitive.E{Key: "path", Value: "$address"}}}}
		group := bson.D{{Key: "$group", Value: bson.D{primitive.E{Key: "_id", Value: "$addressid"}, {Key: "count", Value: bson.D{primitive.E{Key: "$sum", Value: 1}}}}}}
		pointcursor, err := UserCollection.Aggregate(ctx, mongo.Pipeline{matchfilter, unwind, group})
		if err != nil {
			c.IndentedJSON(500, "Internal Server Error")
		}
		var addressinfo []bson.M
		if err = pointcursor.All(ctx, &addressinfo); err != nil {
			panic(err)
		}
		var size int32
		for _, addressno := range addressinfo {
			count := addressno["count"]
			size = count.(int32)
		}
		if size < 2 {
			filter := bson.D{primitive.E{Key: "_id", Value: usertid}}
			update := bson.D{{Key: "$push", Value: bson.D{primitive.E{Key: "address", Value: addresses}}}}
			_, err := UserCollection.UpdateOne(ctx, filter, update)
			if err != nil {
				c.IndentedJSON(500, "Internal Server Error")
			}
		} else {
			c.IndentedJSON(400, "Not Allowed")
		}
		ctx.Done()
	}
}
func EditHomeAddress() gin.HandlerFunc {
	return func(c *gin.Context) {
		userid := c.Query("id")
		if userid == "" {
			c.Header("content-Type", "application/json")
			c.JSON(http.StatusNotFound, gin.H{"Error": "Invalid"})
			c.Abort()
		}
		usertid, err := primitive.ObjectIDFromHex(userid)
		if err != nil {
			c.IndentedJSON(500, "Internal Server Error")
		}
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var editaddress models.Address
		if c.BindJSON(&editaddress); err != nil {
			c.IndentedJSON(http.StatusBadRequest, err.Error())
		}
		filter := bson.D{primitive.E{Key: "_id", Value: usertid}}
		update := bson.D{{Key: "$set", Value: bson.D{primitive.E{Key: "address.0.housename", Value: editaddress.House}, {Key: "address.0.streetname", Value: editaddress.Street}, {Key: "address.0.city", Value: editaddress.City}, {Key: "address.0.pincode", Value: editaddress.Pincode}}}}
		_, err = UserCollection.UpdateOne(ctx, filter, update)
		if err != nil {
			c.IndentedJSON(500, "Something Went Wrong")
			return
		}
		ctx.Done()
		c.IndentedJSON(200, "Sucessfully updated the home address")
	}
}
func EditWorkAddress() gin.HandlerFunc {
	return func(c *gin.Context) {
		userid := c.Query("id")
		if userid == "" {
			c.Header("content-Type", "application/json")
			c.JSON(http.StatusNotFound, gin.H{"Error": "Invalid"})
			c.Abort()
		}
		usertid, err := primitive.ObjectIDFromHex(userid)
		if err != nil {
			c.IndentedJSON(500, "Internal Server Error")
		}
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var editaddress models.Address
		if c.BindJSON(&editaddress); err != nil {
			c.IndentedJSON(http.StatusBadRequest, err.Error())
		}
		filter := bson.D{primitive.E{Key: "_id", Value: usertid}}
		update := bson.D{{Key: "$set", Value: bson.D{primitive.E{Key: "address.1.housename", Value: editaddress.House}, {Key: "address.1.streetname", Value: editaddress.Street}, {Key: "address.1.city", Value: editaddress.City}, {Key: "address.1.pincode", Value: editaddress.Pincode}}}}
		_, err = UserCollection.UpdateOne(ctx, filter, update)
		if err != nil {
			c.IndentedJSON(500, "something went wrong")
			return
		}
		ctx.Done()
		c.IndentedJSON(200, "Successfully Updated")
	}
}
func DeleteAddress() gin.HandlerFunc {
	return func(c *gin.Context) {
		userid := c.Query("id")
		if userid == "" {
			c.Header("content-Type", "application/json")
			c.JSON(http.StatusNotFound, gin.H{"Error": "Invalid search index"})
			c.Abort()
			return
		}
		addresses := make([]models.Address, 0)
		usertid, err := primitive.ObjectIDFromHex(userid)
		if err != nil {
			c.IndentedJSON(500, "Internal Server Error")
		}
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		filter := bson.D{primitive.E{Key: "_id", Value: usertid}}
		update := bson.D{{Key: "$set", Value: bson.D{primitive.E{Key: "address", Value: addresses}}}}
		_, err = UserCollection.UpdateOne(ctx, filter, update)
		if err != nil {
			c.IndentedJSON(404, "wrong command")
			return
		}
		defer cancel()
		ctx.Done()
		c.IndentedJSON(200, "Sucessfully Deleted")

	}
}
