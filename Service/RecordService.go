package Service

import (
	"fmt"
	"restful-api/Infrastructure/Repository"
	"restful-api/Model"
	"restful-api/Shared/Utils"
)

/*
	Returns a list of records, descriptive message, and code for status.
 */
func FilterRecords(request Model.RecordRequest) ([]Model.RecordItem, string, int) {

	if request.StartDate == nil || request.EndDate == nil || request.MinCount == nil || request.MaxCount == nil {
		return nil, "Please provide all the fields (startDate, endDate, minCount and maxCount)", 2
	}

	if *request.MinCount > *request.MaxCount {
		return nil, fmt.Sprintf("MinCount '%d' should be smaller than or equal to MaxCount '%d'",
			*request.MinCount, *request.MaxCount), 3
	}

	if !Utils.ValidateDateFormat(*request.StartDate) {
		return nil, fmt.Sprintf("Start date '%s' is not in the YYYY-MM-DD format", *request.StartDate), 4
	}

	if !Utils.ValidateDateFormat(*request.EndDate) {
		return nil, fmt.Sprintf("End date '%s' is not in the YYYY-MM-DD format", *request.EndDate), 5
	}

	startDate := Utils.ParseDate(*request.StartDate)
	endDate := Utils.ParseDate(*request.EndDate)

	if startDate.Unix() > endDate.Unix() {
		return nil, fmt.Sprintf("StartDate '%s' should be smaller than or equal to EndDate '%s'",
			*request.StartDate, *request.EndDate), 6
	}

	filteredRecords := Repository.FilterByDateAndSum(startDate, endDate, *request.MinCount, *request.MaxCount)

	if len(filteredRecords) == 0 {
		return nil, "Couldn't find any record with given arguments", 7
	}

	return filteredRecords, "Success", 0

}