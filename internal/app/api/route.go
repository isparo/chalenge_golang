package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
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
		}
	}

	err := r.Run(":8080")
	if err != nil {
		glog.Error("can not start service")
	}
}

func pong(c *gin.Context) {
	log.Println("pong handler")
	c.JSON(http.StatusOK, gin.H{"msg": "pong"})
}
