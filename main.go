package main

import (
	"Ptncafe.Golang.Mongo.Test/constant"
	"Ptncafe.Golang.Mongo.Test/provider/mongo_provider"
	"Ptncafe.Golang.Mongo.Test/routes"
	"context"
	"github.com/gin-gonic/gin"
)

func main() {
	ctx:= context.Background()
	shippingDiscountDb,err := mongo_provider.NewMongoProvider(ctx, constant.MongoDbConnectionString, constant.MongoDbNameShippingDiscount)
	if err !=nil{
		panic(err)
	}
	r := gin.Default()
	routes.RegisterRoute(r, shippingDiscountDb)
	r.Run()
}
