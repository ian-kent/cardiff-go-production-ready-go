package main

import (
	"log"
	"net/http"

	"expvar"
)

// curl http://localhost:8080/
// curl http://localhost:8080/debug/vars | jq '.status'

var (
	requests = expvar.NewInt("count")
	status   = expvar.NewMap("status")
)

func main() {
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		requests.Add(1)
		status.Add("200", 1)
		w.WriteHeader(200)
	}))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
