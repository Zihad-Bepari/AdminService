package util

import (
	"encoding/json"
	"net/http"
)

func Senddata(w http.ResponseWriter, data interface{}, statuscode int) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	w.WriteHeader(statuscode)

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}
}

func SendError(w http.ResponseWriter, statuscode int, msg string) {
	w.WriteHeader(statuscode)
	encoder := json.NewEncoder(w)
	encoder.Encode(msg)
}
