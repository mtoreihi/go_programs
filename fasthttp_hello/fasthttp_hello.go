package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/valyala/fasthttp"
)

var (
	addr     = flag.String("addr", ":8080", "TCP address to listen to")
	compress = flag.Bool("compress", false, "Whether to enable transparent response compression")
)

func main() {
	flag.Parse()

	h := requestHandler
	if *compress {
		h = fasthttp.CompressHandler(h)
	}

	if err := fasthttp.ListenAndServe(*addr, h); err != nil {
		log.Fatalf("Error in ListenAndServe: %s", err)
	}
}

func requestHandler(ctx *fasthttp.RequestCtx) {
	if (string(ctx.RequestURI()) == "/test") {
		fmt.Fprintf(ctx, "<html><body>")

		//fmt.Fprintf(ctx, "Request method is %q\n", ctx.Method())
		//fmt.Fprintf(ctx, "RequestURI is %q\n", ctx.RequestURI())
		//fmt.Fprintf(ctx, "Requested path is %q\n", ctx.Path())
		//fmt.Fprintf(ctx, "Host is %q\n", ctx.Host())
		//fmt.Fprintf(ctx, "Query string is %q\n", ctx.QueryArgs())
		//fmt.Fprintf(ctx, "User-Agent is %q\n", ctx.UserAgent())
		fmt.Fprintf(ctx, "Request From: %q<br>", ctx.RemoteAddr())
		fmt.Fprintf(ctx, "Server Time: %s<br>", ctx.ConnTime())
		fmt.Fprintf(ctx, "</body></html>")
		//fmt.Fprintf(ctx, "Request has been started at %s\n", ctx.Time())
		//fmt.Fprintf(ctx, "Serial request number for the current connection is %d\n", ctx.ConnRequestNum())
		

		//fmt.Fprintf(ctx, "Raw request is:\n---CUT---\n%s\n---CUT---", &ctx.Request)

		ctx.SetContentType("text/html; charset=utf8")
	}

	// Set arbitrary headers
	//ctx.Response.Header.Set("X-My-Header", "my-header-value")

	// Set cookies
	/*var c fasthttp.Cookie
	c.SetKey("cookie-name")
	c.SetValue("cookie-value")
	ctx.Response.Header.SetCookie(&c)*/
}