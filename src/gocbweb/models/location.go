// location
package models

import (
	"github.com/QLeelulu/goku"
	"gocbweb"
	"os"
)

type Location struct {
	//	Id      string `json:"Id"`
	//	Type    string `json:"Type"`
	Meta    MyDoc
	City    string `json:"City"`
	State   string `json:"State"`
	Country string `json:"Country"`
}

type LocationQueryRow struct {
	ID    string
	Key   interface{}
	Value Location
	Doc   *Location
}

type LocationQuery struct {
	TotalRows int
	Rows      []LocationQueryRow
}

func GetAllLocations() []Location {
	bucket := GetDB()
	defer bucket.Close()
	var locations []Location
	var query LocationQuery
	err := bucket.ViewCustom(gocbweb.DDOC, "all_location",
		map[string]interface{}{"full_set": true}, &query)
	if err != nil {
		goku.Logger().Logln(err)
		os.Exit(1)
	}
	for _, row := range query.Rows {
		locations = append(locations, row.Value)
	}
	return locations
}
