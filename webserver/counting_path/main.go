package counting_path

import (
	"fmt"
	"log"
	"net/http"
	"sort"
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
	fmt.Fprintf(w, "Welcome, %q, You've called %q %d times\n", r.RemoteAddr, r.URL.Path, counts[r.URL.Path])

	keys := make([]string, 0, len(r.Header))
	for k := range r.Header {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		v := r.Header[k]
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	mu.Unlock()
}
