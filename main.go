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

	r.GET("/hello/:name", func(c *cai.Context) {
		// expect /hello/caijianhao
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.GET("/assets/*filepath", func(c *cai.Context) {
		c.JSON(http.StatusOK, cai.H{"filepath": c.Param("filepath")})
	})

	r.Run(":9999")
}
