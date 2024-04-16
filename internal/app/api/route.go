package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

func (a apiV1) loadRoutes() {
	glog.Info("loading routes")

	r := gin.Default()
	r.GET("/ping", pong)

	v1 := r.Group("/api/v1")
	{
		user := v1.Group("/user")
		{
			user.POST("", a.userHandler.CreateUser)
		}
	}

	//err := r.Run(api.apiConfig.Host + ":" + api.apiConfig.Port)
	err := r.Run(":8080")
	if err != nil {
		glog.Error("can not start service")
	}

}

func pong(c *gin.Context) {
	glog.Info("pong handler")
	c.JSON(http.StatusOK, gin.H{"msg": "pong"})
}
