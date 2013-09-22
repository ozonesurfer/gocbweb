// config
package gocbweb

import (
	"github.com/QLeelulu/goku"
	"path"
	"runtime"
	"time"
)

var Config *goku.ServerConfig = &goku.ServerConfig{
	Addr:           "localhost:8080",
	ReadTimeout:    10 * time.Second,
	WriteTimeout:   10 * time.Second,
	MaxHeaderBytes: 1 << 20,
	//RootDir:        os.Getwd(),
	StaticPath: "static",
	ViewPath:   "views",
	LogLevel:   goku.LOG_LEVEL_LOG,
	Debug:      true,
}

const (
	PASSWORD  = "babcock"
	DDOC      = "dev_gocbweb"
	BUCKET    = "gocbweb"
	POOL      = "127.0.0.1:8091/"
	FULLPOOL  = "http://" + BUCKET + ":" + PASSWORD + "@" + POOL
	BANDTYPE  = "band"
	LOCTYPE   = "location"
	GENRETYPE = "genre"
)

func init() {
	/**
	 * project root dir
	 */
	_, filename, _, _ := runtime.Caller(1)
	Config.RootDir = path.Dir(filename)
}
