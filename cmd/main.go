package main

import (
	"Calendar/interntal/handler"
	"Calendar/interntal/model"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	md := model.Event{Event: "dsfsd", UUID: 432423, UserID: 43254, Date: time.Now()}

	bt, _ := json.Marshal(md)
	fmt.Println(string(bt))
	handler.InitHandlers()
	fmt.Printf("Starting server\n")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
