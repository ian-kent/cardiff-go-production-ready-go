package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// curl http://localhost:8080/

func main() {
	s := &http.Server{
		Addr: ":8080",
	}

	exit := make(chan struct{})
	sigs := make(chan os.Signal, 1)
	go func() {
		<-sigs
		d := time.Now().Add(5 * time.Second)
		ctx, cancel := context.WithDeadline(context.Background(), d)

		defer cancel()

		log.Println("Shutting down server...")
		if err := s.Shutdown(ctx); err != nil {
			log.Fatalf("could not shutdown: %v", err)
		}
		close(exit)
	}()
	signal.Notify(sigs, os.Interrupt)

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

	<-exit
}
