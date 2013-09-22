// band
package controllers

import (
	"fmt"
	"github.com/QLeelulu/goku"

	//	"gocbweb/models"
	"gocbweb"
	"gocbweb/models2"
)

var BandController = goku.Controller("band").
	Get("add", func(ctx *goku.HttpContext) goku.ActionResulter {
	ctx.ViewData["Title"] = "Adding A Band"
	locations := models2.GetAll(gocbweb.LOCTYPE)
	return ctx.View(locations)
	//return ctx.Html("not implemented")
}).Post("verify", func(ctx *goku.HttpContext) goku.ActionResulter {
	ctx.ViewData["Title"] = "Verifying Band"
	name := ctx.Request.FormValue("name")
	loctype := ctx.Request.FormValue("loctype")
	var locationId string
	errorString := "no errors"
	switch loctype {
	case "existing":
		if ctx.Request.FormValue("location_id") == "" {
			errorString = "No location was selected"
		} else {
			locationId = ctx.Request.FormValue("location_id")
		}
		break
	case "new":
		if ctx.Request.FormValue("country") != "" {
			locationId = models2.GenerateId(gocbweb.LOCTYPE)
			location := models2.Location{
				City:    ctx.Request.FormValue("city"),
				State:   ctx.Request.FormValue("state"),
				Country: ctx.Request.FormValue("country")}
			doc := models2.MyDoc{Id: locationId, Type: gocbweb.LOCTYPE, Value: location}
			added, err := models2.AddDoc(doc, gocbweb.LOCTYPE)
			if err != nil {
				errorString = "error on location add: " + err.Error()
			}
			if added == false {
				errorString = "Duplicate id: " + locationId
			}
		} else {
			errorString = "Country is required"
		}
		break
	}
	if errorString == "no errors" {
		var albums []models2.Album
		id := models2.GenerateId(gocbweb.BANDTYPE)
		band := models2.Band{Name: name, LocationId: locationId, Albums: albums}
		doc := models2.MyDoc{Id: id, Type: gocbweb.BANDTYPE, Value: band}
		added, err := models2.AddDoc(doc, gocbweb.BANDTYPE)
		if err != nil {
			errorString = fmt.Sprintf("Band add: %s", err.Error())
		}
		if added == false {
			errorString = "Duplicate id: " + id
		}
	}
	ctx.ViewData["Message"] = errorString
	return ctx.View(nil)
})
