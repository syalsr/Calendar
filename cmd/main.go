package main

import (
	"Calendar/interntal/handler"
	"fmt"
	"log"
	"net/http"
)

func main() {
	handler.InitHandlers()
	fmt.Printf("Starting server\n")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
