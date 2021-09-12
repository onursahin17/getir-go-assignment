package Repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"reflect"
	"restful-api/Infrastructure/Database"
	"restful-api/Model"
	"time"
)

type RecordRepository struct {
	Ctx    context.Context
	Client *mongo.Client
}

var recordRepositoryInstance *RecordRepository = nil

// Creates mongodb connection, initializes an empty record repository instance
func GetRecordRepositoryInstance() *RecordRepository {
	if recordRepositoryInstance == nil {
		ctx, client := Database.InitDbConnection(
			"mongodb+srv://challengeUser:WUMglwNBaydH8Yvu@challenge-xzwqd.mongodb.net/getir-case-study?retryWrites=true")

		recordRepositoryInstance = &RecordRepository{
			Ctx:    ctx,
			Client: client,
		}
	}
	return recordRepositoryInstance
}

// Get a collection (table) based on a given database name and a table name
func getCollectionByName(dbName, tableName string) *mongo.Collection {
	database := GetRecordRepositoryInstance().Client.Database(dbName)
	recordsCollection := database.Collection(tableName)
	return recordsCollection
}

/*
	First, filter records collection by date (obtain the records with date greater
	than the start date and less than the end date). Then, obtain the records with
	total count between min and max count values. Return records that meet both criteria.
 */
func FilterByDateAndSum(startDate time.Time, endDate time.Time, minCount int32, maxCount int32) []Model.RecordItem {

	recordsCollection := getCollectionByName("getir-case-study", "records")

	filterDateCursor, err := recordsCollection.Find(recordRepositoryInstance.Ctx,
		bson.M{"createdAt": bson.M{"$gte": startDate, "$lte": endDate}})
	if err != nil {
		log.Fatal(err)
	}
	var recordsFilteredByDate []bson.M
	if err = filterDateCursor.All(recordRepositoryInstance.Ctx, &recordsFilteredByDate); err != nil {
		log.Fatal(err)
	}

	var recordsFiltered []Model.RecordItem
	for _, record := range recordsFilteredByDate {
		var totalCount int32
		var arr = reflect.ValueOf(record["counts"])
		for i := 0; i < arr.Len(); i++ {
			totalCount += arr.Index(i).Interface().(int32)
		}
		if totalCount >= minCount && totalCount <= maxCount {
			recordItem := Model.RecordItem{
				Key:        reflect.ValueOf(record["key"]).Interface().(string),
				CreatedAt:  reflect.ValueOf(record["createdAt"]).Interface().(primitive.DateTime),
				TotalCount: totalCount,
			}
			recordsFiltered = append(recordsFiltered, recordItem)
		}
	}

	return recordsFiltered

}
