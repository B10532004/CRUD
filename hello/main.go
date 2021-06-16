package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	ConnectMysql()
	r.GET("/users", LoginAuth)
	r.POST("/user", SignUp)
	r.PUT("/user/:id", ChangeProfile)
	r.DELETE("/user/:id", Destroy)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
