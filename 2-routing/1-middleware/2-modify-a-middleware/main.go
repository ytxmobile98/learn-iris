package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/basicauth"
)

func main() {
	users := map[string]string{
		"username": "password",
	}
	auth := basicauth.Default(users)

	app := iris.New()

	app.UseRouter(skipStaticSubdomain(auth))

	app.Get("/", func(ctx iris.Context) {
		ctx.HTML("<h1>Hello world!</h1>")
	})

	app.Listen(":8080")
}

func skipStaticSubdomain(handler iris.Handler) iris.Handler {
	return func(ctx iris.Context) {
		if ctx.Subdomain() == "static" {
			// continue to the next or main handler and exit.
			ctx.Next()
			return
		}

		handler(ctx)
	}
}
