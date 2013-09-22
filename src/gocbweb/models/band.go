// band
package models

import (
	"github.com/QLeelulu/goku"
	"gocbweb"
	"os"
)

type Band struct {
	//	Id         string `json:"Id"`
	//	Type       string `json:"Type"`
	Meta       MyDoc
	Name       string `json:"Name"`
	LocationId string `json:"LocationId"`
	Albums     []Album
}

type Album struct {
	Name    string
	Year    int
	GenreId string
}

type BandQueryRow struct {
	ID    string
	Key   interface{}
	Value Band
	Doc   *Band
}

type BandQuery struct {
	TotalRows int
	Rows      []BandQueryRow
}

func GetAllBands() []Band {
	bucket := GetDB()
	defer bucket.Close()
	var bands []Band
	var query BandQuery
	err := bucket.ViewCustom(gocbweb.DDOC, "all_band",
		map[string]interface{}{"full_set": true}, &query)
	if err != nil {
		goku.Logger().Logln(err)
		os.Exit(1)
	}
	for _, row := range query.Rows {
		bands = append(bands, row.Value)
	}
	return bands
}
