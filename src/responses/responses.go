package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if error := json.NewEncoder(w).Encode(data); error != nil {
		log.Fatal(error)
	}
}

func Error(w http.ResponseWriter, statusCode int, error error) {
	JSON(w, statusCode, struct {
		Error string `json:"erro"`
	}{
		Error: error.Error(),
	})
}
