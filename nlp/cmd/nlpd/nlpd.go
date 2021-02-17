package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/dhiller/go-teach/nlp"
)

func main() {
	// routing

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Post("/api/health", healthHandler)
	r.Get("/api/tokenize", handleTokenize)

	// pre check error before starting service
	err := healthCheck()
	if err != nil {
		log.Fatal(err)
	}

	// running
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func handleTokenize(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	// step 1: get and unmarshall data
	const maxSize = 1 << 20 // 1MB
	data, err := ioutil.ReadAll(io.LimitReader(r.Body, maxSize))
	if err != nil {
		// TODO: (to debate) pass internal error to API?
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	text := string(data)

	// step 2: work
	tokens := nlp.Tokenize(text)

	// step 3: marshal + send
	resp := map[string]interface{}{
		"tokens": tokens,
	}
	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("can't marshal JSON: %s", err)
	}


}

func healthCheck() error {
	return nil
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	err := healthCheck()
	if err == nil {
		fmt.Fprintf(w, "OK")
	} else {
		w.WriteHeader(500)
	}
}
