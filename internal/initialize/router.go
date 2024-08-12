package initialize

import (
	"fmt"
	c "golang-BE/internal/controller"
	"golang-BE/internal/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AA() gin.HandlerFunc {
  return func(c *gin.Context) {
    fmt.Println("Before -->> AA")
    c.Next()
    fmt.Println("After -->> AA")
  }
}

func BB() gin.HandlerFunc {
  return func(c *gin.Context) {
    fmt.Println("Before -->> BB")
    c.Next()
    fmt.Println("After -->> BB")
  }
}

func CC(c *gin.Context) {
  fmt.Println("Before -->> CC")
  c.Next()
  fmt.Println("After -->> CC")
}


func InitRouter() *gin.Engine {

	r := gin.Default()
  //use middleware
  r.Use(middlewares.AuthenticationMiddleware(),BB(), CC)

  v1 := r.Group(("v1/2024"))
  {
    v1.GET("/ping", Pong)
    v1.GET("/user/1", c.NewUserController().GetUserByID)
    // v1.PUT("/ping", Pong)
    // v1.PATCH("/ping", Pong)
    // v1.DELETE("/ping", Pong)
    // v1.HEAD("/ping", Pong)
    // v1.OPTIONS("/ping", Pong)
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