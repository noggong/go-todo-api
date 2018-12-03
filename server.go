package main

import (
	"log"
	"net/http"
)

func main() {
	//port := 4040

	//http.HandleFunc("/", Index)
	//http.HandleFunc("/todos", TodoIndex)


	//log.Printf("Server starting on port %v\n", port)
	//log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))

	router := NewRouter()
	log.Fatal(http.ListenAndServe(":4040", router))
}
