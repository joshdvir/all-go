package main

import (
	iris "gopkg.in/kataras/iris.v5"
)

func init() {
	iris.Config.ReadBufferSize = 1024 * 200
}

func main() {

	iris.OnError(iris.StatusNotFound, func(ctx *iris.Context) {
		ctx.Write("all go!")
		ctx.SetStatusCode(200)
		responseHeaders(ctx, "GET")
	})

	iris.Listen(":9292")
}

func responseHeaders(ctx *iris.Context, methods string) {
	ctx.SetHeader("Access-Control-Allow-Origin", "*")
	ctx.SetHeader("Access-Control-Allow-Methods", methods)
	ctx.SetHeader("Access-Control-Max-Age", "3600")
}
