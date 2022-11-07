package main

import (
	"cai"
	"net/http"
)

func main() {
	r := cai.New()
	r.GET("/", func(c *cai.Context) {
		c.HTML(http.StatusOK, "<h1>Hello cai</h1>")
	})
	r.GET("/hello", func(c *cai.Context) {
		// expect /hello?name=caiktutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *cai.Context) {
		c.JSON(http.StatusOK, cai.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":9999")
}
