package Utils

import (
	"log"
	"regexp"
	"time"
)

func ValidateDateFormat(date string) bool {
	/*
		Accept dates that have YYYY-MM-DD format
		yyyy -> starts with 19 or 20
		mm -> between 00 and 12
		dd -> between 00 and 31
	*/
	dateRegex := regexp.MustCompile("((19|20)\\d\\d)-(0?[1-9]|1[012])-(0?[1-9]|[12][0-9]|3[01])")
	return dateRegex.MatchString(date)
}

func ParseDate(date string) time.Time {
	parsedDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		log.Fatal(err)
	}
	return parsedDate
}
