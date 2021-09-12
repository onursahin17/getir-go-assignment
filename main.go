package main

import (
	"log"
	"net/http"
	"os"
	RecordController "restful-api/Controller"
	"restful-api/Infrastructure/Repository"
)

func main() {
	/*
	Initialize the repository instances, create connections
	- Record Repository creates connection with given mongodb
	- In Memory Repository creates an empty in-memory instance
	 */
	Repository.GetRecordRepositoryInstance()
	Repository.GetInMemoryRepositoryInstance()
	/// ****** ///
	HandleRequests()
}

func HandleRequests() {
	http.HandleFunc("/filtered-records", RecordController.FilteredRecords)
	http.HandleFunc("/in-memory", RecordController.GetValue)
	http.HandleFunc("/in-memory/set", RecordController.SetValue)

	log.Fatal(http.ListenAndServe(":" + os.Getenv("PORT"), nil))
}
