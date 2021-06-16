package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("login.html")
	//設定靜態資源的讀取
	r.Static("/assets", "custom.css")
	r.GET("/login", LoginPage)
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	r.POST("/login", LoginAuth)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
