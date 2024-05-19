package main

import "github.com/kataras/iris/v12"

func main() {
	app := iris.New()

	app.Use(iris.Compression)
	app.Use(checkGzip)

	app.Get("/", rootHandler)
	app.Post("/data", dataHandler)

	app.Listen(":8080")
}

func checkGzip(ctx iris.Context) {
	header := ctx.ResponseWriter().Header()
	headerName := "x-client-supports-gzip"

	if ctx.ClientSupportsEncoding("gzip") {
		header.Set(headerName, "true")
	} else {
		header.Set(headerName, "false")
	}

	ctx.Next()
}

func rootHandler(ctx iris.Context) {
	ctx.HTML("<h1>GZIP compression</h1>")
}

func dataHandler(ctx iris.Context) {
	response := iris.Map{"name": "kataras"}
	ctx.JSON(response)
}
