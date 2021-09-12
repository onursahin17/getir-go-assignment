package Model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RecordItem struct {
	Key        string             `json:"key"`
	CreatedAt  primitive.DateTime `json:"createdAt"`
	TotalCount int32              `json:"totalCount"`
}

type RecordResponse struct {
	Code    int          `json:"code"`
	Message string       `json:"msg"`
	Records []RecordItem `json:"records"`
}

type InMemoryResponse struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}