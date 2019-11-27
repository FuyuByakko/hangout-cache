package main

import (
	"encoding/json"
	"fmt"
	"hangout-cache/cache"
	"log"
	"net/http"
)

var (
	serverPort = ":9000"
	version    = "v1"
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
	log.Fatal(http.ListenAndServe(serverPort, nil))
}

func reqHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		checkCache(w, r)
	} else if r.Method == "POST" {
		writeToCache(w, r)
	}
}

func writeToCache(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method)
	fmt.Fprintf(w, "POST endpoint reached")
	//optional:
	keys := r.URL.Query()
	hashNr := keys.Get("key") // < returns empty string if not found.
	cache.Add(hashNr, "Hello!!")
}

func checkCache(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method)
	keys, ok := r.URL.Query()["key"]
	if !ok || len(keys[0]) < 1 {
		log.Println("URL Param 'key' is missing")
		return
	}
	//Right now keys is an array with our hash#
	hashNr := keys[0]

	//check the cache for the hash key
	response := cache.Get(hashNr)
	if response != "" {
		log.Printf("Found Hash Nr %s in the cache, returning associated Value", hashNr)
	}
	json.NewEncoder(w).Encode(response)

	// fmt.Fprintf(w, "THIS IS A RESPONSE! Hello from port %s!\nHash is %s", serverPort, hashNr)
}

func defaultGreet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
	Welcome to the caching service.
	Please use GET 'api/%s/contents/?key='hash' to retrieve data for a specific hash.
	Use same endpoint with POST to create a new entry.`, version)
}
