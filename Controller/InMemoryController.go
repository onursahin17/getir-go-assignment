package Controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"restful-api/Infrastructure/Repository"
	"restful-api/Model"
	InMemoryService "restful-api/Service"
)

/*
	Endpoint to get key-value pair
	It calls InMemoryService function with the provided key in URL as an argument.
*/
func GetValue(w http.ResponseWriter, r *http.Request){

	key := r.URL.Query().Get("key")
	if key == "" {
		http.Error(w, "Error: Missing key name in query string!", http.StatusBadRequest)
		return
	}
	value, err := InMemoryService.GetValue(key)
	if err == Repository.KeyNotFoundErr {
		http.Error(w, fmt.Sprintf("%s", err), http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, fmt.Sprintf("Error ocurred while getting value from database: %s", err),
			http.StatusInternalServerError)
		return
	}

	response := Model.InMemoryResponse{
		Key:   key,
		Value: *value,
	}
	json.NewEncoder(w).Encode(response)

}

/*
	Endpoint to set key-value pair
	It decodes the request and pass it to InMemoryRequest struct.
	Then, calls RecordService function with the request struct as an argument.
	Handle errors and print descriptive error messages (ex: if key is nil, value is nil)
*/
func SetValue(w http.ResponseWriter, r *http.Request){

	var inMemoryRequest Model.InMemoryRequest

	err := json.NewDecoder(r.Body).Decode(&inMemoryRequest)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid request structure: %s", err), http.StatusBadRequest)
		return
	}
	if inMemoryRequest.Key == nil {
		http.Error(w, "Error: Please provide the key", http.StatusBadRequest)
		return
	}
	if inMemoryRequest.Value == nil {
		http.Error(w, "Error: Please provide the value", http.StatusBadRequest)
		return
	}
	InMemoryService.SetValue(inMemoryRequest)
	// Echoing the request
	json.NewEncoder(w).Encode(inMemoryRequest)
}
