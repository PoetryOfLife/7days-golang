package main

import (
	"net/http"
	"zed"
)

func main() {
	//z := zed.New()
	//z.GET("/", func(c *zed.Context) {
	//	c.HTML(http.StatusOK, "<h1>helle</h1>")
	//})
	//z.GET("/hello", func(c *zed.Context) {
	//	c.String(http.StatusOK, "hello %s,you're at %s\n", c.Query("name"), c.Path)
	//})
	//z.GET("/hello/:name", func(c *zed.Context) {
	//	c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	//})
	//
	//z.GET("/assets/*filepath", func(c *zed.Context) {
	//	c.JSON(http.StatusOK, zed.H{"filepath": c.Param("filepath")})
	//})
	//z.Run(":9999")

	z := zed.New()
	z.GET("/index", func(c *zed.Context) {
		c.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})
	v1 := z.Group("/v1")
	{
		v1.GET("/", func(c *zed.Context) {
			c.HTML(http.StatusOK, "<h1>Hello zed</h1>")
		})

		v1.GET("/hello", func(c *zed.Context) {
			// expect /hello?name=zedktutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}
	v2 := z.Group("/v2")
	{
		v2.GET("/hello/:name", func(c *zed.Context) {
			// expect /hello/zedktutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
		v2.POST("/login", func(c *zed.Context) {
			c.JSON(http.StatusOK, zed.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})

	}

	z.Run(":9999")
}
