// base
package models

import (
	"github.com/QLeelulu/goku"
	"github.com/couchbaselabs/go-couchbase"
	"gocbweb"
	"os"
)

type MyDoc struct {
	Id   string `json:"Id"`
	Type string `json:"Type"`
}

func GetDB() *couchbase.Bucket {
	b, err := couchbase.GetBucket(gocbweb.FULLPOOL, "default", gocbweb.BUCKET)
	if err != nil {
		goku.Logger().Logln(err)
		os.Exit(1)
	}
	return b
}
