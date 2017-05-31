package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// curl http://localhost:8080/

func main() {
	s := http.Server{
		Addr:              ":8080",
		ReadTimeout:       1 * time.Second,
		WriteTimeout:      1 * time.Second,
		ReadHeaderTimeout: 1 * time.Second,
		IdleTimeout:       1 * time.Second,
		MaxHeaderBytes:    0,
	}

	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(200)
		for i := 0; i < 5; i++ {
			time.Sleep(1)
			w.Write([]byte(fmt.Sprintf("%d", i)))
		}
	}))

	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
