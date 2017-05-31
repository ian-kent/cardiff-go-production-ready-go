package main

import (
	"log"
	"net/http"
)

// curl http://localhost:8080/

func main() {
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(200)
	}))

	// Don't do this!
	http.Handle("/healthcheck", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if !checkDBConnection() {
			w.WriteHeader(500)
			return
		}

		w.WriteHeader(200)
		return
	}))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func checkDBConnection() bool {
	for i := 0; i < 10; i++ {
		log.Println("Burning CPU")
		var v int
		for v < 1000000000 {
			v++
		}
	}
	// DB is ok!
	return true
}
