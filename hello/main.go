package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	//全域middleware
	r := gin.New()
	r.Use(gin.Logger())
	ConnectMysql()
	ConnectRedis()
	//By Group設定middleware
	g1 := r.Group("/v1").Use(middleware1)
	g1.GET("/users", UserList)          //R
	g1.POST("/user", SignUp)            //C
	g1.PUT("/user/:id", ChangePassword) //U
	g1.DELETE("/user/:id", Destroy)     //D
	g2 := r.Group("/v2").Use((middleware2))
	g2.GET("/usertimes", CountAPI) //R
	r.Run()                        // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func middleware1(c *gin.Context) {
	c.Next()
	AddAPI()
}

func middleware2(c *gin.Context) {
	c.Next()
}
