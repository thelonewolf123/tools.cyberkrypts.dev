package controllers

import (
	"github.com/gin-gonic/gin"
	"tools.cyberkrypts.dev/interfaces"
	"tools.cyberkrypts.dev/templates/pages"
	"tools.cyberkrypts.dev/utils"
)

type HomeController struct{}

var landingPageInfo = []interfaces.LandingPageInfo{{
	Title:       "Youtube",
	Description: "Download Youtube videos with ease",
	URL:         "/youtube",
	Background:  "blue",
},
	{
		Title:       "Shortener",
		Description: "Shorten and share URLs with ease",
		URL:         "/shortener",
		Background:  "green",
	},
	{
		Title:       "Json Formatter",
		Description: "Format and validate JSON with ease",
		URL:         "/json-formatter",
		Background:  "yellow",
	},
	// {
	// 	Title:       "Talk Swipe",
	// 	Description: "Swipe right for interesting conversations",
	// 	URL:         "/talk-swipe",
	// 	Background:  "purple",
	// }, {
	// 	Title:       "Send Files",
	// 	Description: "Send files to anyone with a link",
	// 	URL:         "/send-files",
	// 	Background:  "red",
	// },
}

func (c HomeController) Index(ctx *gin.Context) {
	utils.RenderTemplate(200, ctx, pages.Home(landingPageInfo))
}
