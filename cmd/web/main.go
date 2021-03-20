package main

import (
	"github.com/ahojo/hello-world-web/pkg/handlers"
	"net/http"
)


const portNumber = ":8000"


func main() {
	//fmt.Println("Hello World")
	http.HandleFunc("/about", handlers.About)

	http.HandleFunc("/", handlers.Home)


	_ = http.ListenAndServe(portNumber, nil)
}


