package handler

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Message string `json:"error"`
}

func WriteJSON(w http.ResponseWriter, data any, statusCode int) {
	if data == nil {
		w.WriteHeader(statusCode)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	bodyBytes, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if _, err = w.Write(bodyBytes); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func WriteError(w http.ResponseWriter, err error, statusCode int) {
	WriteJSON(w, ErrorResponse{Message: err.Error()}, statusCode)
}

func ReadJSON(r *http.Request, data any) error {
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(data); err != nil {
		return err
	}
	return nil
}
