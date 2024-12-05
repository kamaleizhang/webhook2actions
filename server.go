package hookact

import (
	"fmt"
	"io"
	"net/http"

	"hookact/actions/hugo"
)

func handleRoot(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading body", http.StatusInternalServerError)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			http.Error(w, "Error closing body", http.StatusInternalServerError)
		}
	}(r.Body)
	fmt.Printf("Received request: %s\n", body)
	_, _ = fmt.Fprintf(w, "Hello, World! This is a simple HTTP server in Go.")
}

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleRoot)
	mux.HandleFunc("/hugo", hugo.HandleHook)
	return mux
}

func StartServer(addr string, mux *http.ServeMux) error {
	return http.ListenAndServe(addr, mux)
}
