package handler

import "net/http"

func InitHandlers() {
	http.Handle("/create_event", CreateEvent)
	http.Handle("/update_event", UpdateEvent)
	http.Handle("/delete_event", DeleteEvent)
	http.Handle("/events_for_day", EventsForDay)
	http.Handle("/events_for_week", EventsForWeek)
	http.Handle("/events_for_month", EventsForMonth)
}

func CreateEvent(writter http.ResponseWriter, request *http.Request) {

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
