package main

import (
	"net/http"

	"github.com/kataras/iris/v12"
)

func routerWrapper(w http.ResponseWriter, r *http.Request, router http.HandlerFunc) {
	if r.URL.Path == "/" {
		w.Write([]byte("#1 .WrapRouter\n"))
		/* Note for new Gophers:
		   If we Write anything here on an error resource in the raw
		   `net/http` wrapper like this one, then the response writer will
		   automatically send a `200` OK status code (when we first write).
		   Any error handler executed after this will not fire as expected.
		   Also, when `w.WriteHeader` is called you can NOT change the
		   status code later on.

		   In Iris Handlers, if you write before the status code has been
		   set, then it will also automatically send the 200 OK status
		   code which then cannot be changed later. However, if we call
		   `ctx.StatusCode` inside an Iris Handler without writing any
		   content, then we can change the status code later on. When you
		   need to change that behaviour, you must start the handler with
		   a `ctx.Record` call.
		*/
	}

	// Continue by executing the Iris Router and let it do its job.
	router(w, r)
}

func routerMiddleware(ctx iris.Context) {
	if ctx.Path() == "/" {
		ctx.WriteString("#2 .UseRouter\n")
		// The same caveat described in routerWrapper applies here as well.
	}

	ctx.Next()
}

func globalMiddleware(ctx iris.Context) {
	ctx.WriteString("#3 .UseGlobal\n")
	ctx.Next()
}

func useMiddleware(ctx iris.Context) {
	ctx.WriteString("#4 .Use\n")
	ctx.Next()
}

func errorMiddleware(ctx iris.Context) {
	ctx.WriteString("#3 .UseError\n")
	ctx.Next()
}
