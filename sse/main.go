package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	writerChans := map[string]chan string{}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	mux.HandleFunc("/sse", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/event-stream")

		writerChan := make(chan string)
		writerChans[r.RemoteAddr] = writerChan

		for {
			msg := <-writerChan
			_, err := w.Write([]byte("data: " + msg + "\n\n"))
			if err != nil {
				log.Printf("could not write to response writer: %v", err)
				delete(writerChans, r.RemoteAddr)
				return
			}
			w.(http.Flusher).Flush()
		}
	})

	mux.HandleFunc("/send", func(w http.ResponseWriter, r *http.Request) {
		for _, writerChan := range writerChans {
			writerChan <- "Test"
		}
	})

	err := http.ListenAndServe("127.0.0.1:9000", mux)
	if err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}
