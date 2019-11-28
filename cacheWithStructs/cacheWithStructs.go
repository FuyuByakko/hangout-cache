package cacheWithStructs

//this is a version of cache that uses Structs as a blueprint to parse the JSON object and store it

import (
	"hangout-cache/structs"
	"log"
)

//examples can be found in structs/testStructs
var testEvents = structs.TestResponse().Events
var testHotels = structs.TestResponse().Hotels
var EventsAndHotes = structs.EventsAndHotels{testEvents, testHotels}

var Cache = make(map[string]structs.EventsAndHotels)

// func CreateCache() *cache {
// 	var Cache = make(map[string]structs.EventsAndHotels)
// 	return &cache
// }

func Get(key string) (structs.EventsAndHotels, error) {
	return Cache[key], nil
}

func Add(k string, s structs.EventsAndHotels) {
	Cache[k] = s
	check := Cache[k]
	log.Printf("%s\n", Cache[k])
}
