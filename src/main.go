package main

import (
	"fmt"
	"log"
	"net/http"

	controller "github.com/notionals-lab/pixel/src/controller"
	// "github.com/gorilla/websocket"
)

func main() {

	http.HandleFunc("/", controller.GetSocketHandler)

	// route
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello!")
	})

	http.HandleFunc("/api/pixels", controller.GetPixelHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
