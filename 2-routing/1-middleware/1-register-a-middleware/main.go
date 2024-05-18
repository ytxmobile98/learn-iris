package main

import (
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()

	app.WrapRouter(routerWrapper)
	app.UseRouter(routerMiddleware)
	app.UseGlobal(globalMiddleware)
	app.Use(useMiddleware)
	app.UseError(errorMiddleware)

	// app.Done(done)
	// app.DoneGlobal(doneGlobal)

	// Adding a OnErrorCode(iris.StatusNotFound) causes `.UseGlobal`
	// to be fired on 404 pages without this,
	// only `UseError` will be called, and thus should
	// be used for error pages.
	app.OnErrorCode(iris.StatusNotFound, notFoundHandler)

	app.Get("/", mainHandler)

	app.Listen(":8080")
}

func mainHandler(ctx iris.Context) {
	ctx.WriteString("Main Handler")
}

func notFoundHandler(ctx iris.Context) {
	ctx.WriteString("404 Error Handler")
}
