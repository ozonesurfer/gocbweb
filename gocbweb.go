// gocbweb
package main

import (
	"github.com/QLeelulu/goku"
	"gocbweb"
	"gocbweb/controllers"
	"log"
)

func main() {
	rt := &goku.RouteTable{Routes: gocbweb.Routes}
	s := goku.CreateServer(rt, nil, gocbweb.Config)
	goku.Logger().Logln("Server start on", s.Addr)
	log.Fatal(s.ListenAndServe())

}

var home = controllers.HomeController
var band = controllers.BandController
var album = controllers.AlbumController
