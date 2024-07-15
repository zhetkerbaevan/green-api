package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)


func ParseJSON(r *http.Request, payload any) error {
	//Check if body is empty
	if r.Body == nil {
		return fmt.Errorf("MISSING REQUEST BODY")
	}
	return json.NewDecoder(r.Body).Decode(payload) //Create new decoder which will read data from r.Body, then decode reads data from json and decode it to payload
}

func WriteJSON(w http.ResponseWriter, v any) error {
	//Add header (json) for response
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v) //Create new encoder which will write data to w, then encode will write data from v to json
}

