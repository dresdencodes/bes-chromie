package captureserve

import (
	"log"
	"net/http"
)

// next content to serve
var serveContent = make(chan string, 10)

// next html
func NextHTML(html string) {
	serveContent <- html
}

// get url
func Run() {


	// validate server already running
	server := &http.Server{Addr: ":11111"}

	// handle func 
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(<-serveContent))
	})

	log.Println("Serving on http://localhost:11111")
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal(err)
	}

}