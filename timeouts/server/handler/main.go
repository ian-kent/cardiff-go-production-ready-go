package main

import (
	"log"
	"net/http"
	"time"
)

// curl http://localhost:8080/

func main() {
	http.Handle("/", http.TimeoutHandler(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		time.Sleep(2 * time.Second)
		w.WriteHeader(200)
		w.Write([]byte("¯\\_(ツ)_/¯\n"))
	}), 1*time.Second, "request timed out"))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
