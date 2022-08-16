package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/sunist-c/toy-framework/core/http"
)

/**
result of this example:
INFO[0000] [hello-world] GET: /hello ~> main.main.func2  operation="bind middleware" when="2022-08-16:17-54-12.CST" where="github.com/sunist-c/toy-framework/core/http.(*Service).bind-$ProjectPath/toy-framework/core/http/service.go:82"
INFO[0000] [hello-world] GET: /hello ~> main.main.func3  operation="bind middleware" when="2022-08-16:17-54-12.CST" where="github.com/sunist-c/toy-framework/core/http.(*Service).bind-$ProjectPath/toy-framework/core/http/service.go:82"
INFO[0000] [hello-world] GET: /hello ~> main.main.func4  operation="bind middleware" when="2022-08-16:17-54-12.CST" where="github.com/sunist-c/toy-framework/core/http.(*Service).bind-$ProjectPath/toy-framework/core/http/service.go:82"
INFO[0000] [hello-world] GET: /hello ~> main.main.func5  operation="bind middleware" when="2022-08-16:17-54-12.CST" where="github.com/sunist-c/toy-framework/core/http.(*Service).bind-$ProjectPath/toy-framework/core/http/service.go:82"
INFO[0000] [hello-world] GET: /hello -> main.main.func1  operation="bind handler" when="2022-08-16:17-54-12.CST" where="github.com/sunist-c/toy-framework/core/http.(*Service).bind-$ProjectPath/toy-framework/core/http/service.go:84"
INFO[0000] server started at 0.0.0.0:8080                operation="start up server" when="2022-08-16:17-54-12.CST" where="github.com/sunist-c/toy-framework/core/http.(*Engine).Serve-$ProjectPath/toy-framework/core/http/engine.go:46"

run shell command:
curl 127.0.0.1:8080/hello

result of response
"hello, world"

console output:
Hello, world!
Func1
Func2
Func3
Uses: 8.75µs
*/

func main() {
	e := http.NewEngine()
	e.RegisterService(&http.Service{
		ServiceName: "hello-world",
		URL:         "hello",
		Father:      nil,
		Info:        nil,
		Method:      http.GET,
		Handler: func(context *gin.Context) {
			context.JSON(200, "hello, world")
		},
		Middlewares: []gin.HandlerFunc{
			func(context *gin.Context) {
				fmt.Println("Hello, world!")
				start := time.Now()
				context.Next()
				fmt.Println("Uses:", time.Now().Sub(start))
			},
			func(context *gin.Context) {
				fmt.Println("Func1")
				context.Next()
			},
			func(context *gin.Context) {
				fmt.Println("Func2")
				context.Next()
			},
			func(context *gin.Context) {
				fmt.Println("Func3")
				context.Next()
			},
		},
	})
	e.Serve("0.0.0.0:8080")
	defer e.Close()
	time.Sleep(time.Second * 60)
}
