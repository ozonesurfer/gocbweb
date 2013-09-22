// album
package controllers

import (
	"fmt"
	"github.com/QLeelulu/goku"
	"gocbweb"
	"gocbweb/models2"
	"strconv"
)

var AlbumController = goku.Controller("album").
	Get("index", func(ctx *goku.HttpContext) goku.ActionResulter {
	bandId := ctx.RouteData.Params["id"]
	bandDoc := models2.GetDoc(bandId)
	bandValue := bandDoc.Value.(map[string]interface{})
	ctx.ViewData["Title"] = bandValue["Name"].(string)
	//	ctx.ViewData["Name"] = bandValue["Name"].(string)
	//	ctx.ViewData["Id"] = bandDoc.Id
	//	albums := bandValue["Albums"].([]models2.Album)
	//	albums := bandDoc.GetAlbums()
	return ctx.View(bandDoc)
}).Get("add", func(ctx *goku.HttpContext) goku.ActionResulter {
	ctx.ViewData["Title"] = "Add Album"
	ctx.ViewData["Id"] = ctx.RouteData.Params["id"]
	genres := models2.GetAll(gocbweb.GENRETYPE)
	return ctx.View(genres)
}).Post("verify", func(ctx *goku.HttpContext) goku.ActionResulter {
	ctx.ViewData["Title"] = "Verifying Album"
	rawId := ctx.RouteData.Params["id"]
	name := ctx.Request.FormValue("name")
	yearString := ctx.Request.FormValue("year")
	year, _ := strconv.Atoi(yearString)
	genretype := ctx.Request.FormValue("genretype")
	var genreId string
	errorString := "no errors"
	switch genretype {
	case "existing":
		if ctx.Request.FormValue("genre_id") == "" {
			errorString = "No genre was selected"
		} else {
			genreId = ctx.Request.FormValue("genre_id")
		}
		break
	case "new":
		if ctx.Request.FormValue("genre_name") != "" {
			genreId = models2.GenerateId(gocbweb.GENRETYPE)
			genre := models2.Genre{Name: ctx.Request.FormValue("genre_name")}
			doc := models2.MyDoc{Id: genreId, Type: gocbweb.GENRETYPE,
				Value: genre}
			added, err := models2.AddDoc(doc, gocbweb.GENRETYPE)
			if err != nil {
				errorString = fmt.Sprintf("Genre: %s", err.Error())
			}
			if added == false {
				errorString = "Duplicate id: " + genreId
			}
		} else {
			errorString = "Genre name is required"
		}
		break
	}

	if errorString == "no errors" {
		bandId := rawId
		bandDoc := models2.GetDoc(bandId)
		album := models2.Album{Name: name, Year: year, GenreId: genreId}
		err := bandDoc.AddAlbum(album)
		if err != nil {
			errorString = fmt.Sprintf("Album: %s", err.Error())
		}
	}
	ctx.ViewData["Message"] = errorString
	ctx.ViewData["Id"] = rawId
	return ctx.View(nil)
})
