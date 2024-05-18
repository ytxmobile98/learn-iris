package main

import (
	"github.com/kataras/iris/v12"
)

const (
	key   = "key"
	value = "Hello world!"
)

func main() {
	app := iris.New()

	app.Get("/", myMiddleware, myHandler)

	app.Listen(":8080")
}

func myMiddleware(ctx iris.Context) {
	ctx.Values().Set(key, value)

	ctx.Next() // must be used inside a middleware
}

func myHandler(ctx iris.Context) {
	value := ctx.Values().GetString(key)
	ctx.HTML("<h1>%s</h1>", value)
}
