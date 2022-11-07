package main

import (
	"cai"
	"net/http"
)

func main() {
	r := cai.New()
	r.GET("/index", func(c *cai.Context) {
		c.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})
	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *cai.Context) {
			c.HTML(http.StatusOK, "<h1>Hello cai</h1>")
		})

		v1.GET("/hello", func(c *cai.Context) {
			// expect /hello?name=caiktutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}
	v2 := r.Group("/v2")
	{
		v2.GET("/hello/:name", func(c *cai.Context) {
			// expect /hello/caiktutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
		v2.POST("/login", func(c *cai.Context) {
			c.JSON(http.StatusOK, cai.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})

	}

	r.Run(":9999")
}
