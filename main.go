package main

import (
	"os"
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()

	port := os.Getenv("PORT")
	if port == "" {
		port = "9292"
	}

	app.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
		ctx.Writef("all go!")
		ctx.StatusCode(200)
		responseHeaders(ctx, "GET,POST")
	})

	app.Run(iris.Addr(":" + port))
}

func responseHeaders(ctx iris.Context, methods string) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Methods", methods)
	ctx.Header("Access-Control-Max-Age", "3600")
}
