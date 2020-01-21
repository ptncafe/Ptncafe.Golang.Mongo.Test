package main

import (
	"Ptncafe.Golang.Mongo.Test/constant"
	"Ptncafe.Golang.Mongo.Test/provider/mongo_provider"
	"context"
	"github.com/gin-gonic/gin"
)

func main() {
	ctx:= context.Background()
	shippingDiscountDb,err := mongo_provider.NewMongoProvider(ctx, constant.MongoDbConnectionStringectionString, constant.MongoDbConnectionStringNameShippingDiscount)
	if err !=nil{
		panic(err)
	}
	r := gin.Default()
	//r.GET("queue", controller.QueueController)
	r.Run()
}
