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

const layout = "2006-01-02"

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
		writter.WriteHeader(http.StatusMethodNotAllowed)
		writter.Write([]byte("{error: bad request}"))
		return
	}

	var event model.Event

	body, err := io.ReadAll(request.Body)
	if err != nil {
		writter.Write([]byte(`{"error": "bad request - incorrect body"}`))
		return
	}

	js := make(map[string]string)
	err = json.Unmarshal(body, &js)
	if err != nil {
		log.Println(err)
		return
	}

	event.Event = js["event"]

	event.Date, err = time.Parse(layout, js["date"])
	if err != nil {
		log.Println(err)
		return
	}

	event.UserID, err = strconv.Atoi(js["user_id"])
	if err != nil {
		log.Println(err)
		return
	}

	event.UUID = api.Calendar.LastUUID + 1
	api.Calendar.LastUUID = event.UUID

	api.Calendar.CreateEvent(&event)

	writter.WriteHeader(http.StatusOK)
	writter.Write([]byte("{result: success"))

	log.Println(event)
}

func UpdateEvent(writter http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		writter.WriteHeader(http.StatusMethodNotAllowed)
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
	err = json.Unmarshal(body, &js)
	if err != nil {
		log.Println(err)
		return
	}

	event.Event = js["event"]

	event.Date, err = time.Parse(layout, js["date"])
	if err != nil {
		log.Println(err)
		return
	}

	event.UserID, err = strconv.Atoi(js["user_id"])
	if err != nil {
		log.Println(err)
		return
	}

	event.UUID, err = strconv.Atoi(js["UUID"])
	if err != nil {
		log.Println(err)
		return
	}

	api.Calendar.UpdateEvent(&event)

	writter.WriteHeader(http.StatusOK)
	writter.Write([]byte("{result: success"))

	log.Println(event)
}

func DeleteEvent(writter http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		writter.WriteHeader(http.StatusMethodNotAllowed)
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
	err = json.Unmarshal(body, &js)
	if err != nil {
		log.Println(err)
		return
	}

	event.UUID, err = strconv.Atoi(js["UUID"])
	if err != nil {
		log.Println(err)
		return
	}

	api.Calendar.DeleteEvent(event.UUID)

	writter.WriteHeader(http.StatusOK)
	writter.Write([]byte("{result: success"))
}

func EventsForDay(writter http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodPost {
		writter.WriteHeader(http.StatusMethodNotAllowed)
		writter.Write([]byte("{error: bad request}"))
		return
	}

	userID, err := strconv.Atoi(request.URL.Query().Get("user_id"))
	if err != nil {
		log.Println(err)
		return
	}

	date, err := time.Parse(layout, request.URL.Query().Get("date"))
	if err != nil {
		log.Println(err)
		return
	}

	events := api.Calendar.GetEventPerDay(date, userID)
	for _, event := range events {
		js, err := json.MarshalIndent(event, "", "\t")
		if err != nil {
			log.Println(err)
		}
		writter.Write(js)
	}
	writter.WriteHeader(http.StatusOK)
}

func EventsForWeek(writter http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodPost {
		writter.WriteHeader(http.StatusMethodNotAllowed)
		writter.Write([]byte("{error: bad request}"))
		return
	}

	userID, err := strconv.Atoi(request.URL.Query().Get("user_id"))
	if err != nil {
		log.Println(err)
		return
	}

	date, err := time.Parse(layout, request.URL.Query().Get("date"))
	if err != nil {
		log.Println(err)
		return
	}

	events := api.Calendar.GetEventPerDay(date, userID)
	for _, event := range events {
		js, err := json.MarshalIndent(event, "", "\t")
		if err != nil {
			log.Println(err)
		}
		writter.Write(js)
	}
	writter.WriteHeader(http.StatusOK)
}

func EventsForMonth(writter http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodPost {
		writter.WriteHeader(http.StatusMethodNotAllowed)
		writter.Write([]byte("{error: bad request}"))
		return
	}

	userID, err := strconv.Atoi(request.URL.Query().Get("user_id"))
	if err != nil {
		log.Println(err)
		return
	}

	date, err := time.Parse(layout, request.URL.Query().Get("date"))
	if err != nil {
		log.Println(err)
		return
	}

	events := api.Calendar.GetEventPerDay(date, userID)
	for _, event := range events {
		js, err := json.MarshalIndent(event, "", "\t")
		if err != nil {
			log.Println(err)
		}
		writter.Write(js)
	}
	writter.WriteHeader(http.StatusOK)
}
