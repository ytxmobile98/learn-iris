package main

import "github.com/kataras/iris/v12"

func main() {
	app := iris.New()

	app.SetExecutionRules(iris.ExecutionRules{
		Begin: iris.ExecutionOptions{Force: true},
		Main:  iris.ExecutionOptions{Force: true},
		Done:  iris.ExecutionOptions{Force: true},
	})

	app.Use(getMiddleware(1), getMiddleware(2))

	app.Get("/", func(ctx iris.Context) {
		ctx.HTML("<h1>Hello world!</h1>")
	})

	app.Listen(":8080")
}

func getMiddleware(i uint) iris.Handler {
	return func(ctx iris.Context) {
		ctx.HTML("<h1>Middleware: %d</h1>", i)
	}
}
