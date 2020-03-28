package main

import (
	"github.com/afifialaa/handlers"
	"fmt"
	"net/http"

)

type Status struct{
	msg string
}

func main(){

	//routes
	http.HandleFunc("/user/signup", handlers.SignupHandle)
	http.HandleFunc("/user/signin", handlers.LoginHandle)
	http.HandleFunc("/user/test", handlers.LoginHandle)

	//listening for requests
	fmt.Println("server is running")
	http.ListenAndServe(":8080", nil)

}
