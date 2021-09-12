package Service

import (
	"restful-api/Infrastructure/Repository"
	"restful-api/Model"
)

// Gets value based on provided key
func GetValue(key string) (*string, error) {
	return Repository.GetInMemoryRepositoryInstance().Get(key)
}

// Sets key-value pair based on provided key and value
func SetValue(request Model.InMemoryRequest) {
	Repository.GetInMemoryRepositoryInstance().Set(*request.Key, *request.Value)
}
