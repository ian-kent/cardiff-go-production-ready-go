package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// curl http://localhost:8080/

func main() {
	s := &http.Server{
		Addr: ":8080",
	}

	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(200)
		for i := 0; i < 5; i++ {
			time.Sleep(1 * time.Second)
			w.Write([]byte(fmt.Sprintf("%d\n", i)))
			w.(http.Flusher).Flush()
		}
	}))

	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
