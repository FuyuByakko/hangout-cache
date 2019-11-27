package cache

import (
	"hangout-cache/structs"
	"log"
)

var testEvents = structs.TestResponse().Events
var testHotels = structs.TestResponse().Hotels
var EventsAndHotes = structs.EventsAndHotels{testEvents, testHotels}

var Cache = make(map[string]structs.EventsAndHotels)

// func CreateCache() *cache {
// 	var Cache = make(map[string]structs.EventsAndHotels)
// 	return &cache
// }

func Get(key string) (structs.EventsAndHotels, error) {
	Cache["test"] = EventsAndHotes
	// if Cache[key] != ""
	return Cache[key], nil
}

func Add(k string, s structs.EventsAndHotels) {
	log.Println(k)
	log.Println(s)
	Cache[k] = s
}

// func (ev *EventsAndHotels) Add(event Event, hotel Hotel) {
// 	ev.Events = append(ev.Events, event);
// 	ev.Hotels = append(ev.Hotels, hotel);
// }
