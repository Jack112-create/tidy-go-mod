package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	_ "github.com/davecgh/go-spew/spew"
)

func main() {
	http.HandleFunc("/routes", GetRoutes)

	fmt.Println("Server listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}

func GetRoutes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	resp, err := http.Get("https://galway-bus.apis.ie/api")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sb := string(body)
	err = json.NewEncoder(w).Encode(sb)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	log.Printf(sb)
}
