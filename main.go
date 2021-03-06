package main

import (
	"os"
	"fmt"
	"encoding/json"
	"net/http/httputil"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
)

func PrettyPrint(v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
					fmt.Println(string(b))
	}
	return
}

func main() {
	app := iris.New()

	port := os.Getenv("PORT")
	if port == "" {
		port = "9292"
	}

	customLogger := logger.New(logger.Config{
		// Status displays status code
		Status: true,
		// IP displays request's remote address
		IP: true,
		// Method displays the http method
		Method: true,
		// Path displays the request path
		Path: true,
		// Query appends the url query to the Path.
		Query: true,

		// Columns: true,

		// if !empty then its contents derives from `ctx.Values().Get("logger_message")
		// will be added to the logs.
		MessageContextKeys: []string{"logger_message"},

		// if !empty then its contents derives from `ctx.GetHeader("User-Agent")
		MessageHeaderKeys: []string{"User-Agent"},
	})

	app.Use(customLogger)

	app.OnErrorCode(iris.StatusNotFound, customLogger, func(ctx iris.Context) {
		ctx.StatusCode(iris.StatusOK)
		ctx.Writef("all go!")
		fmt.Println("Request: \n")
		responseHeaders(ctx, "GET,POST")
	})

	app.Post("/service/v7/haas_sync", func(ctx iris.Context) {
		ctx.Writef("{\"SyncResponse\":{\"status\":\"success\",\"Users\":0,\"CamerasHash\":\"\",\"Cameras\":[]}}")
		ctx.StatusCode(iris.StatusOK)
		responseHeaders(ctx, "GET,POST")
	})

	app.Run(iris.Addr(":" + port))
}

func responseHeaders(ctx iris.Context, methods string) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Methods", methods)
	ctx.Header("Access-Control-Max-Age", "3600")
	requestDump, err := httputil.DumpRequest(ctx.Request(), true)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(requestDump))
}
