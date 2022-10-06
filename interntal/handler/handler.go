package handler

import (
	"Calendar/interntal/api"
	"Calendar/interntal/model"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

func InitHandlers() {
	http.HandleFunc("/create_event", CreateEvent)
	http.HandleFunc("/update_event", UpdateEvent)
	http.HandleFunc("/delete_event", DeleteEvent)
	http.HandleFunc("/events_for_day", EventsForDay)
	http.HandleFunc("/events_for_week", EventsForWeek)
	http.HandleFunc("/events_for_month", EventsForMonth)
}

func CreateEvent(writter http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		writter.Write([]byte("{error: bad request}"))
		return
	}
	var event model.Event
	body, err := io.ReadAll(request.Body)
	if err != nil {
		writter.WriteHeader(http.StatusBadRequest)
		writter.Write([]byte(`{"error": "bad request - incorrect body"}`))
		return
	}

	js := make(map[string]string)
	json.Unmarshal(body, &js)

	event.Event = js["event"]
	event.Date, _ = time.Parse("2000-10-05", js["date"])
	event.UserID, _ = strconv.Atoi(js["user_id"])
	event.UUID = api.Calendar.LastUUID + 1
	api.Calendar.LastUUID = event.UUID

	api.Calendar.CreateEvent(&event)
	writter.Write([]byte("{result: success"))

	log.Println(event)
}

func UpdateEvent(writter http.ResponseWriter, request *http.Request) {

}

func DeleteEvent(writter http.ResponseWriter, request *http.Request) {

}

func EventsForDay(writter http.ResponseWriter, request *http.Request) {

}

func EventsForWeek(writter http.ResponseWriter, request *http.Request) {

}

func EventsForMonth(writter http.ResponseWriter, request *http.Request) {

}
