package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Go Docker Running!")

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(struct {
			Status string `json:"status"`
		}{
			Status: "pong",
		})
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}
