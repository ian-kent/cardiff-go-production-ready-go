package main

import (
	"context"
	"log"
	"net/http"
	"time"
)

// curl http://localhost:8080/

func main() {
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		someExpensiveOp(req.Context())
		w.WriteHeader(200)
	}))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func someExpensiveOp(ctx context.Context) {
	var stop bool
	go func() {
		<-ctx.Done()
		log.Println("stopping")
		stop = true
	}()
	for i := 0; i < 10; i++ {
		if stop {
			log.Printf("stopped")
			break
		}
		log.Printf("sleeping %d...", i)
		time.Sleep(1 * time.Second)
	}
	log.Printf("done")
}
