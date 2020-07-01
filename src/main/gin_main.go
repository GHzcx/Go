package main

import "github.com/gin-gonic/gin"

type data struct {
	Name string `name`
	Age int `age`
}
func main() {
	r := gin.Default()
	d := data{"haha", 15}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code": 0,
			"message": "success",
			"data" : d,
		})
	})
	r.Run("0.0.0.0:8000") // 监听并在 0.0.0.0:8080 上启动服务
}
