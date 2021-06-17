package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	ConnectMysql()
	ConnectRedis()
	r.GET("/users", UserList)          //R
	r.POST("/user", SignUp)            //C
	r.PUT("/user/:id", ChangePassword) //U
	r.DELETE("/user/:id", Destroy)     //D
	r.Run()                            // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
