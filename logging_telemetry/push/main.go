package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type output struct {
	Timestamp time.Time   `json:"timestamp"`
	Event     string      `json:"event"`
	Data      interface{} `json:"data"`
}

func writeToKafka(o output) {
	// FIXME write to a real kafka topic
	b, _ := json.Marshal(o)
	fmt.Fprintf(os.Stdout, "%s\n", b)
}

func main() {
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		writeToKafka(output{time.Now(), "request", map[string]interface{}{"status": 200}})
		w.WriteHeader(200)
	}))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
