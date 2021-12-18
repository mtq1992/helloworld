package main

import (
	"gee"
)

func main() {
	r := gee.New()
	r.GET("/", func(ctx *gee.Context) {
		ctx.HTML(200, "<h1>hello world</h1>")
	})

	r.GET("/hello", func(ctx *gee.Context) {
		ctx.String(200, "hello &s, you at %s\n", ctx.Query("name"), ctx.Path)
	})

	r.POST("/login", func(ctx *gee.Context) {
		ctx.JSON(200, gee.H{
			"username": ctx.PostForm("username"),
			"password": ctx.PostForm("password"),
		})
	})

	r.Run(":9999")
}
