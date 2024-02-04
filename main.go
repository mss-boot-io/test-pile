package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	port = flag.Int("port", 80, "port to listen on")
	path = flag.String("path", "/", "path to listen on")
)

func main() {
	flag.Parse()
	r := gin.Default()
	r.POST(*path, func(c *gin.Context) {
		rb, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}
		println(string(rb))
		c.Status(200)
	})
	err := r.Run(fmt.Sprintf(":%d", *port)) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err != nil {
		fmt.Println(err)
	}
}
