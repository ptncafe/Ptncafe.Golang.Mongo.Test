package routes

import (
	"Ptncafe.Golang.Mongo.Test/controller"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func RegisterRoute(r *gin.Engine, shippingDiscountDb *mongo.Database){
	r.GET("ping", controller.PingController(shippingDiscountDb))
}