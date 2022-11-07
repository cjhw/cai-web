package main

import (
	"cai"
	"net/http"
)

func main() {
	r := cai.Default()
	r.GET("/", func(c *cai.Context) {
		c.String(http.StatusOK, "Hello caiktutu\n")
	})
	// index out of range for testing Recovery()
	r.GET("/panic", func(c *cai.Context) {
		names := []string{"caiktutu"}
		c.String(http.StatusOK, names[100])
	})

	r.Run(":9999")
}
