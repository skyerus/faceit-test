package event

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/skyerus/faceit-test/pkg/user"
)

// List of URLs to hit with create/update/delete events
type eventListeners struct {
	create []string
	update []string
	delete []string
}

// Example of a service subscribed to all events
const example1Service string = "http://localhost:8181"

// Example of a service subscribed to just create events
const example2Service string = "http://localhost:8182"

var appEventListeners eventListeners = eventListeners{
	create: []string{example1Service, example2Service},
	update: []string{example1Service},
	delete: []string{example1Service},
}

// BroadcastCreateEvent ...
func BroadcastCreateEvent(u user.User) {
	if os.Getenv("NO_EVENT_BROADCASTS") == "true" {
		return
	}
	byteData, err := json.Marshal(u)
	if err != nil {
		log.Fatal(err)
	}
	b := bytes.NewBuffer(byteData)
	for _, host := range appEventListeners.create {
		req, err := http.NewRequest("POST", host+"/users", b)
		if err != nil {
			log.Fatal(err)
		}
		err = sendRequest(req)
		if err != nil {
			log.Println(err)
		}
	}
}

// BroadcastUpdateEvent ...
func BroadcastUpdateEvent(u user.User) {
	if os.Getenv("NO_EVENT_BROADCASTS") == "true" {
		return
	}
	byteData, err := json.Marshal(u)
	if err != nil {
		log.Fatal(err)
	}
	b := bytes.NewBuffer(byteData)
	for _, host := range appEventListeners.update {
		req, err := http.NewRequest("PUT", host+"/users/"+strconv.Itoa(u.ID), b)
		if err != nil {
			log.Fatal(err)
		}
		err = sendRequest(req)
		if err != nil {
			log.Println(err)
		}
	}
}

// BroadcastDeleteEvent ...
func BroadcastDeleteEvent(ID int) {
	if os.Getenv("NO_EVENT_BROADCASTS") == "true" {
		return
	}
	for _, host := range appEventListeners.delete {
		req, err := http.NewRequest("DELETE", host+"/users/"+strconv.Itoa(ID), nil)
		if err != nil {
			log.Fatal(err)
		}
		err = sendRequest(req)
		if err != nil {
			log.Println(err)
		}
	}
}

func sendRequest(r *http.Request) error {
	client := &http.Client{}
	_, err := client.Do(r)
	return err
}
