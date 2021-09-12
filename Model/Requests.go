package Model

type RecordRequest struct {
	StartDate *string `json:"startDate"`
	EndDate   *string `json:"endDate"`
	MinCount  *int32  `json:"minCount"`
	MaxCount  *int32  `json:"maxCount"`
}

type InMemoryRequest struct {
	Key   *string `json:"key"`
	Value *string `json:"value"`
}