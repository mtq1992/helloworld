package main

import (
	"gee"
	"net/http"
)

func main() {
	r := gee.New()
	r.GET("/", func(ctx *gee.Context) {
		ctx.HTML(200, "<h1>hello world</h1>")
	})

	r.GET("/hello", func(ctx *gee.Context) {
		ctx.String(200, "hello %s, you at %s\n", ctx.Query("name"), ctx.Path)
	})

	r.GET("/hello/:name", func(ctx *gee.Context) {
		// expect /hello/geektutu
		ctx.String(http.StatusOK, "hello %s, you're at %s\n", ctx.Param("name"), ctx.Path)
	})

	r.POST("/login", func(ctx *gee.Context) {
		ctx.JSON(200, gee.H{
			"username": ctx.PostForm("username"),
			"password": ctx.PostForm("password"),
		})
	})

	r.GET("/assets/*filepath", func(ctx *gee.Context) {
		ctx.JSON(http.StatusOK, gee.H{"filepath": ctx.Param("filepath")})
	})

	r.Run(":9999")
}
