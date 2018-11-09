package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	router.POST("/user", CreateUser)

	router.POST("/user/:user_name", Login)

	return router
}

func main() {
	r := RegisterHandlers()
	fmt.Printf("port is :8000,server runing!")
	http.ListenAndServe(":8000", r)

}
