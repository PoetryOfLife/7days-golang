package main

import (
	"net/http"
	"zed"
)

func main() {
	z := zed.New()
	z.GET("/", func(c *zed.Context) {
		c.HTML(http.StatusOK, "<h1>helle</h1>")
	})
	z.GET("/hello", func(c *zed.Context) {
		c.String(http.StatusOK, "hello %s,you're at %s\n", c.Query("name"), c.Path)
	})
	z.GET("/hello/:name", func(c *zed.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	z.GET("/assets/*filepath", func(c *zed.Context) {
		c.JSON(http.StatusOK, zed.H{"filepath": c.Param("filepath")})
	})
	z.Run(":9999")
}
