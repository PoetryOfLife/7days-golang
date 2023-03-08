package main

import (
	"fmt"
	"net/http"
	"zed"
)

func main() {
	z := zed.New()
	z.GET("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "URL.Path = %q\n", request.URL.Path)
	})
	z.GET("/hello", func(writer http.ResponseWriter, request *http.Request) {
		for k, v := range request.Header {
			fmt.Fprintf(writer, "Header[%q] = %q\n", k, v)
		}
	})
	z.Run(":9999")
}
