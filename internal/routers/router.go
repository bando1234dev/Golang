package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouters() *gin.Engine {

	r := gin.Default()

  v1 := r.Group(("v1/2024"))
  {
    v1.GET("/ping", Pong)
    v1.PUT("/ping", Pong)
    v1.PATCH("/ping", Pong)
    v1.DELETE("/ping", Pong)
    v1.HEAD("/ping", Pong)
    v1.OPTIONS("/ping", Pong)
  }

	return r
}


func Pong(c *gin.Context) {
  name := c.DefaultQuery("name","bando1234")
  // c.ShouldBindJSON()
  uid := c.Query("uid")
  c.JSON(http.StatusOK, gin.H{
    "message": "pong...ping" + name,
    "uid": uid,
    "users": []string{"prx","edg","g3"},
  })
}