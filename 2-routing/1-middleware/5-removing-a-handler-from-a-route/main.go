package main

import "github.com/kataras/iris/v12"

func main() {
	app := iris.New()

	// Register the middleware to all matched routes.
	app.Use(middleware)

	// Handlers = middleware, other
	app.Get("/", index)

	// Handlers = other
	app.Get("/other", other).RemoveHandler(middleware)

	app.Listen(":8080")
}

func middleware(ctx iris.Context) {
	ctx.HTML("<h1>middleware</h1>")

	ctx.Next()
}

func index(ctx iris.Context) {
	ctx.HTML("<h1>Hello world!</h1>")
}

func other(ctx iris.Context) {
	ctx.HTML("<h1>/other</h1>")
}
