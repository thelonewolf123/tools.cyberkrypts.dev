package controllers

import (
	"github.com/gin-gonic/gin"
	"tools.cyberkrypts.dev/interfaces"
	"tools.cyberkrypts.dev/templates/pages"
	"tools.cyberkrypts.dev/utils"
)

type HomeController struct{}

var landingPageInfo = []interfaces.LandingPageInfo{{
	Title:           "Youtube",
	Description:     "Download Youtube videos with ease",
	URL:             "/youtube",
	Background:      "bg-blue-500",
	BackgroundHover: "hover:bg-blue-700",
},
	{
		Title:           "Shortener",
		Description:     "Shorten and share URLs with ease",
		URL:             "/shortener",
		Background:      "bg-green-500",
		BackgroundHover: "hover:bg-green-700",
	}, {
		Title:           "Talk Swipe",
		Description:     "Swipe right for interesting conversations",
		URL:             "/talk-swipe",
		Background:      "bg-purple-500",
		BackgroundHover: "hover:bg-purple-700",
	},
}

func (c HomeController) Index(ctx *gin.Context) {
	utils.RenderTemplate(200, ctx, pages.Home(landingPageInfo))
}
