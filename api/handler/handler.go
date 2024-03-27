package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var userName string

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Respond with the user's name
		jsonResponse := map[string]string{"message": fmt.Sprintf("Hello, %s!", userName)}
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(jsonResponse)
		if err != nil {
			fmt.Println("Error encoding JSON response:", err)
			return
		}
		return
	}
	// Handle other requests
	http.NotFound(w, r)
}
