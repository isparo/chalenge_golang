package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (a apiV1) loadRoutes() {
	log.Println("loading routes")

	r := gin.Default()
	r.GET("/ping", pong)

	v1 := r.Group("/api/v1")
	{
		user := v1.Group("/user")
		{
			user.POST("", a.userHandler.CreateUser)
			user.POST("/login", a.userHandler.LogIn)
		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err := r.Run(":8080")
	if err != nil {
		glog.Error("can not start service")
	}
}

func pong(c *gin.Context) {
	log.Println("pong handler")
	c.JSON(http.StatusOK, gin.H{"msg": "pong"})
}
