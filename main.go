package main

import (
	"client-http-1/controllers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/posts", controllers.Index)
	//http.HandleFunc("/post/create", controllers.Create)
	http.HandleFunc("/post/store", controllers.Store)
	//http.HandleFunc("/post/delete", controllers.Delete)

	log.Print("Server started on: http://localhost:8181")
	log.Fatal(http.ListenAndServe(":8181", nil))

}
