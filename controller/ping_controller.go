package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"net/http"
	"time"
)

func PingController(c *gin.Context, shippingDiscountDb *mongo.Database) {
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	err := shippingDiscountDb.Client().Ping(ctx, readpref.Primary())
	if err != nil{
		logrus.Errorf("PingController %v", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": shippingDiscountDb.Name(), "status": http.StatusOK})
	return
}