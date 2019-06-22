package main

import (
	"encoding/json"
	"net/http"
)

// API handler for reversing a string
func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		var data Data
		err := decoder.Decode(&data)
		if err != nil {
			http.Error(w, "Bad Request",
				http.StatusBadRequest)
		}

		// Call to reverse method
		data.reverse()

		res, err := json.Marshal(data)
		if err != nil {
			http.Error(w, "Some error occurred.",
				http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
