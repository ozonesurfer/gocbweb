// base
package models2

import (
	"fmt"
	"github.com/QLeelulu/goku"
	"github.com/couchbaselabs/go-couchbase"
	"gocbweb"
	"os"
	"strconv"
)

type MyDoc struct {
	Id    string `json:"Id"`
	Type  string `json:"Type"`
	Value interface{}
}

type Band struct {
	Name       string `json:"Name"`
	LocationId string `json:"LocationId"`
	Albums     []Album
}

type Album struct {
	Name    string
	Year    int
	GenreId string
}

type Genre struct {
	Name string `json:"Name"`
}

type Location struct {
	City    string `json:"City"`
	State   string `json:"State"`
	Country string `json:"Country"`
}

func GetDB() *couchbase.Bucket {
	b, err := couchbase.GetBucket(gocbweb.FULLPOOL, "default", gocbweb.BUCKET)
	if err != nil {
		goku.Logger().Logln(err)
		os.Exit(1)
	}
	return b
}
func GenerateId(docType string) string {
	bucket := GetDB()
	defer bucket.Close()
	mapString := "all_" + docType
	results, _ := bucket.View(gocbweb.DDOC, mapString,
		map[string]interface{}{"full_set": true})
	count := len(results.Rows)
	id := docType + "-" + strconv.Itoa(count+1)
	return id
}

func GetAll(docType string) []MyDoc {
	bucket := GetDB()
	defer bucket.Close()
	viewString := "all_" + docType
	results, err := bucket.View(gocbweb.DDOC, viewString,
		map[string]interface{}{"full_set": true, "stale": false})
	if err != nil {
		goku.Logger().Logln(err)
		os.Exit(1)
	}
	var docs []MyDoc
	var doc MyDoc

	for _, resultDoc := range results.Rows {
		bucket.Get(resultDoc.ID, &doc)
		docs = append(docs, doc)
	}
	return docs
}

func AddDoc(doc MyDoc, docType string) (bool, error) {
	bucket := GetDB()
	defer bucket.Close()
	doc.Type = docType
	//	added, err := bucket.Add(doc.Id, 0, doc)
	added := true
	err := bucket.Write(doc.Id, 0, 0, doc, couchbase.Persist)
	return added, err
}

func GetDoc(id string) MyDoc {
	bucket := GetDB()
	defer bucket.Close()
	var doc MyDoc
	bucket.Get(id, &doc)
	return doc
}

func GetGenreName(id string) string {
	bucket := GetDB()
	defer bucket.Close()
	var genreDoc MyDoc
	bucket.Get(id, &genreDoc)
	genre := genreDoc.Value.(map[string]interface{})
	name := genre["Name"].(string)
	return name
}

func GetBandsByGenre(id string) []MyDoc {
	bucket := GetDB()
	defer bucket.Close()
	results, err := bucket.View(gocbweb.DDOC, "by_genre",
		map[string]interface{}{"key": id, "full_set": true, "reduce": false})
	if err != nil {
		goku.Logger().Logln(err)
		os.Exit(1)
	}
	var x []MyDoc
	found := false
	for _, row := range results.Rows {
		for _, already := range x {
			if already.Id == row.ID {
				found = true
				break
			}
		}
		if found == false {
			var band MyDoc
			bucket.Get(row.ID, &band)
			x = append(x, band)
		}
		found = false
	}
	return x
}

func (this *MyDoc) LocToString() string {
	var result, cityString, stateString string
	thisDoc := this.Value.(map[string]interface{})
	me := Location{City: thisDoc["City"].(string),
		State:   thisDoc["State"].(string),
		Country: thisDoc["Country"].(string)}
	if me.City != "" {
		cityString = me.City
	} else {
		cityString = "(city)"
	}
	if me.State != "" {
		stateString = me.State
	} else {
		stateString = "(state)"
	}
	result = fmt.Sprintf("%s, %s %s", cityString, stateString, me.Country)
	return result
}

//This only works with Band objects
func (this *MyDoc) GetLocation() string {
	b := GetDB()
	defer b.Close()
	me := this.Value.(map[string]interface{})
	locId := me["LocationId"].(string)
	var locDoc MyDoc
	b.Get(locId, &locDoc)
	locString := locDoc.LocToString()
	return locString
}

type BandDoc struct {
	Id    string
	Type  string
	Value Band
}

func (this *MyDoc) AddAlbum(album Album) error {
	b := GetDB()
	defer b.Close()
	q := this.Value.(map[string]interface{})
	var err error
	if q["Albums"] == nil {
		var albums []Album
		albums = append(albums, album)
		band := Band{Name: q["Name"].(string),
			LocationId: q["LocationId"].(string), Albums: albums}
		doc := MyDoc{Id: this.Id, Type: this.Type, Value: band}
		err = b.Set(doc.Id, 0, doc)
	} else {
		z := q["Albums"].([]interface{})

		var a []Album
		for _, c := range z {
			x := c.(map[string]interface{})

			y := Album{Name: x["Name"].(string),
				Year:    int(x["Year"].(float64)),
				GenreId: x["GenreId"].(string)}
			a = append(a, y)
		}
		a = append(a, album)
		band := Band{Name: q["Name"].(string),
			LocationId: q["LocationId"].(string), Albums: a}
		doc := MyDoc{Id: this.Id, Type: this.Type, Value: band}
		err = b.Set(doc.Id, 0, doc)

	}

	return err
}

func (this *Album) GetGenreName() string {
	bucket := GetDB()
	defer bucket.Close()
	var doc MyDoc
	bucket.Get(this.GenreId, &doc)
	value := doc.Value.(map[string]interface{})
	result := value["Name"].(string)
	//	fmt.Println("Name =", result)
	return result
}

func (this MyDoc) GetAlbums() []Album {
	q := this.Value.(map[string]interface{})
	z := q["Albums"].([]interface{})

	var a []Album
	for _, c := range z {
		x := c.(map[string]interface{})

		y := Album{Name: x["Name"].(string),
			Year:    int(x["Year"].(float64)),
			GenreId: x["GenreId"].(string)}
		a = append(a, y)
	}
	return a
}
