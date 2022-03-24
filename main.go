package main

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/dingtalk/webhook1/send", func(c *gin.Context) {
		var data interface{}
		err := json.NewDecoder(c.Request.Body).Decode(&data)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}
		rb, _ := json.Marshal(data)
		println(string(rb))
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8060") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
