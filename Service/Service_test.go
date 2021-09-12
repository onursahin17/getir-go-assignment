package Service

import (
	"github.com/stretchr/testify/assert"
	"restful-api/Model"
	"testing"
)

/*
	Unit tests for the service functions
 */
func TestFilterRecords(t *testing.T) {

	/*
		TEST1: Request with missing fields as input
		Expected: nil records array, descriptive message, code 2
	 */
	startDate := "2016-01-26"
	endDate := "2018-02-02"

	model := Model.RecordRequest{
		StartDate: &startDate,
		EndDate:   &endDate,
	}
	recordItem, msg, code := FilterRecords(model)

	assert.Nil(t, recordItem)
	assert.Equal(t, "Please provide all the fields (startDate, endDate, minCount and maxCount)", msg)
	assert.Equal(t, 2, code)

	/*
		TEST2: Request with minimum count bigger than maximum count as input
		Expected: nil records array, descriptive message, code 3
	*/
	var minCount int32 = 3100
	var maxCount int32 = 3000
	model = Model.RecordRequest{
		StartDate: &startDate,
		EndDate:   &endDate,
		MinCount:  &minCount,
		MaxCount:  &maxCount,
	}
	recordItem, msg, code = FilterRecords(model)

	assert.Nil(t, recordItem)
	assert.Equal(t, "MinCount '3100' should be smaller than or equal to MaxCount '3000'", msg)
	assert.Equal(t, 3, code)

	/*
		TEST3: Request with invalid start date format as input
		Expected: nil records array, descriptive message, code 4
	*/
	startDate = "201-01-26"
	minCount = 2900
	model = Model.RecordRequest{
		StartDate: &startDate,
		EndDate:   &endDate,
		MinCount:  &minCount,
		MaxCount:  &maxCount,
	}
	recordItem, msg, code = FilterRecords(model)

	assert.Nil(t, recordItem)
	assert.Equal(t, "Start date '201-01-26' is not in the YYYY-MM-DD format", msg)
	assert.Equal(t, 4, code)

	/*
		TEST4: Request with invalid end date format as input
		Expected: nil records array, descriptive message, code 5
	*/
	startDate = "2016-01-26"
	endDate = "20180202"
	model = Model.RecordRequest{
		StartDate: &startDate,
		EndDate:   &endDate,
		MinCount:  &minCount,
		MaxCount:  &maxCount,
	}
	recordItem, msg, code = FilterRecords(model)

	assert.Nil(t, recordItem)
	assert.Equal(t, "End date '20180202' is not in the YYYY-MM-DD format", msg)
	assert.Equal(t, 5, code)

	/*
		TEST5: Request with start date bigger than end date as input
		Expected: nil records array, descriptive message, code 6
	*/
	startDate = "2020-01-26"
	endDate = "2018-02-02"
	model = Model.RecordRequest{
		StartDate: &startDate,
		EndDate:   &endDate,
		MinCount:  &minCount,
		MaxCount:  &maxCount,
	}
	recordItem, msg, code = FilterRecords(model)

	assert.Nil(t, recordItem)
	assert.Equal(t, "StartDate '2020-01-26' should be smaller than or equal to EndDate '2018-02-02'", msg)
	assert.Equal(t, 6, code)

	/*
		TEST6: Request with adjusted fields (ex: close intervals or large dates),
		so that there will be an empty records array after filtering.
		Note: Not unsuccessful but it is better to inform user with a descriptive message
		Expected: nil records array, descriptive message, code 7
	*/
	startDate = "2018-02-02"
	model = Model.RecordRequest{
		StartDate: &startDate,
		EndDate:   &endDate,
		MinCount:  &minCount,
		MaxCount:  &maxCount,
	}
	recordItem, msg, code = FilterRecords(model)

	assert.Nil(t, recordItem)
	assert.Equal(t,  "Couldn't find any record with given arguments", msg)
	assert.Equal(t, 7, code)

	/*
		TEST7: OK request
		Expected: non-empty records array, success message, code 0
	*/
	startDate = "2016-01-26"
	model = Model.RecordRequest{
		StartDate: &startDate,
		EndDate:   &endDate,
		MinCount:  &minCount,
		MaxCount:  &maxCount,
	}
	recordItem, msg, code = FilterRecords(model)

	assert.NotNil(t, recordItem)
	assert.Equal(t,  "Success", msg)
	assert.Equal(t, 0, code)
}
