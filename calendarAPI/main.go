package main

import (
	"fmt"
	"net/http"

	_ "calendarAPI/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

// @contact.name   Hanks Tsai
// @contact.url    https://github.com/HanksJCTsai
// @title Swagger Demo
// @version 1.0
// @description calendarAPP API.
// @host 127.0.0.1:8080
func main() {
	http.HandleFunc("/hello", HelloHandler)
	http.HandleFunc("/swagger/", httpSwagger.WrapHandler)

	http.ListenAndServe(":8080", nil)
}

// @Tags Hello
// @Produce plain
// @Param name query string false "user name"
// @Success 200 {string} string "ok"
// @Router /hello [get]
func HelloHandler(rw http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	content := fmt.Sprintf("hello, %s", name)
	fmt.Fprint(rw, content)
}
