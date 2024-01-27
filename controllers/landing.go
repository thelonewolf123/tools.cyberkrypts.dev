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
	}, {
		Title:       "Talk Swipe",
		Description: "Swipe right for interesting conversations",
		URL:         "/talk-swipe",
		Background:  "purple",
	},
}

func (c HomeController) Index(ctx *gin.Context) {
	utils.RenderTemplate(200, ctx, pages.Home(landingPageInfo))
}
