package main

import (
	"encoding/json"
	"fmt"
	"hangout-cache/cache"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var (
	version = "v1"
)

func main() {
	fmt.Println("hello")
	StartServer()

	// ln, err := http.ListenAndServec("tcp", ":9000")
	// if err != nil {
	// 	panic(err)
	// }
	// for {
	// 	conn, err := ln.Accept()
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Println("handle connection", conn)
	// }

}

func StartServer() {
	fmt.Println("Setting Up Server")
	http.HandleFunc("/", defaultGreet)
	http.HandleFunc("/api/v1/contents/", reqHandler)
	serverPort := ":" + os.Getenv("PORT")
	if serverPort == ":" {
		serverPort = ":9000"
	}
	log.Fatal(http.ListenAndServe(serverPort, nil))
}

//default Greeting returned by the server when entering the main page
func defaultGreet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
	Welcome to the caching service.
	Please use GET 'api/%s/contents/?key='hash' to retrieve data for a specific hash.
	Use same endpoint with POST to create a new entry.`, version)
}

func reqHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		checkCache(w, r)
	} else if r.Method == "POST" {
		writeToCache(w, r)
	}
}

func writeToCache(w http.ResponseWriter, r *http.Request) {
	keys := r.URL.Query()
	hashNr := keys.Get("key") // < returns empty string if not found.
	log.Printf("Received %s Request. Hash is: %s.", r.Method, hashNr)

	//read the Body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	//transform it to a string
	text := string(body)

	//add the received body under the received key into the cache
	cache.Add(hashNr, text)

	//In case of using cacheWithStructs, use below Body parser with Struct
	// var data structs.EventsAndHotels
	// log.Println(r.Body)
	// err := json.NewDecoder(r.Body).Decode(&data)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }
	// cache.Add(hashNr, data)
}

func checkCache(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["key"]
	if !ok || len(keys[0]) < 1 {
		log.Println("URL Param 'key' is missing")
		return
	}
	//Right now keys is an array with our hash#
	hashNr := keys[0]
	log.Printf("Received %s Request. Hash is: %s.", r.Method, hashNr)

	//check the cache for the hash key and send the contents
	response := cache.Get(hashNr)
	json.NewEncoder(w).Encode(response)
}
