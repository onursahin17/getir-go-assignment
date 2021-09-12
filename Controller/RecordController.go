package Controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"restful-api/Model"
	RecordService "restful-api/Service"
)

/*
	Endpoint to fetch data from mongodb.
	It decodes the request and pass it to RecordRequest struct.
	Then, calls RecordService function with the request struct as an argument.
 */
func FilteredRecords(w http.ResponseWriter, r *http.Request){

	var recordRequest Model.RecordRequest
	err := json.NewDecoder(r.Body).Decode(&recordRequest)
	if err != nil {
		response := Model.RecordResponse{
			Code: 1,
			Message: fmt.Sprintf("Invalid request structure: %s", err),
			Records: nil,
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	filteredRecords, message, code := RecordService.FilterRecords(recordRequest)
	response := Model.RecordResponse{Code: code, Message: message, Records: filteredRecords}
	json.NewEncoder(w).Encode(response)
}