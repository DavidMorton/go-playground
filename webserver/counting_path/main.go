package counting_path

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var counts = make(map[string]int)
var mu sync.Mutex

func Run_webserver() {

	fmt.Print("Starting web server...")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	counts[r.URL.Path]++
	fmt.Fprintf(w, "You've called %q %d times", r.URL.Path, counts[r.URL.Path])
	mu.Unlock()
}
