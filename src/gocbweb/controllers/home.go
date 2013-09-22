// home
package controllers

import (
	"fmt"
	"github.com/QLeelulu/goku"
	"gocbweb"
	"gocbweb/models2"
)

var HomeController = goku.Controller("home").
	Get("index", func(ctx *goku.HttpContext) goku.ActionResulter {
	ctx.ViewData["Title"] = "CD Catalog Site"
	bands := models2.GetAll(gocbweb.BANDTYPE)
	return ctx.View(bands)
}).Get("genrelist", func(ctx *goku.HttpContext) goku.ActionResulter {
	ctx.ViewData["Title"] = "List of Genres"
	genres := models2.GetAll(gocbweb.GENRETYPE)
	return ctx.View(genres)
}).Get("bygenre", func(ctx *goku.HttpContext) goku.ActionResulter {
	genreId := ctx.RouteData.Params["id"]
	genreName := models2.GetGenreName(genreId)
	ctx.ViewData["Title"] = fmt.Sprintf("%s Albums", genreName)
	ctx.ViewData["GenreId"] = genreId
	bands := models2.GetBandsByGenre(genreId)
	return ctx.View(bands)
})
