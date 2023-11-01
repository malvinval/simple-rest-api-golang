package main

import (
	"encoding/json"
	"net/http"
)

// crafting the JSON Response format
type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// create a new struct for API response
	var Data Response
	Data.Status = "active"
	Data.Message = "Welcome to our API!"

	// returns the JSON encoding of Data
	jsonEncodedData, err := json.Marshal(Data)

	// JSON encoding error handling
	if err != nil {
		panic(err)
	}

	// set Content-Type header to application/json bcs our API returns JSON
	w.Header().Set("Content-Type", "application/json")

	// OPTIONALLY, u can set the status in response header
	w.WriteHeader(http.StatusOK)

	// return the JSON
	w.Write(jsonEncodedData)
}
