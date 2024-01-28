package controllers

import (
	"fmt"

	"tools.cyberkrypts.dev/db"
	"tools.cyberkrypts.dev/env"
	"tools.cyberkrypts.dev/templates/components"
	"tools.cyberkrypts.dev/templates/pages"
	"tools.cyberkrypts.dev/utils"

	"github.com/gin-gonic/gin"
)

type ShortenerController struct{}

func (sc ShortenerController) Index(ctx *gin.Context) {
	utils.RenderTemplate(200, ctx, pages.ShortenerIndex())
}

func (sc ShortenerController) Generate(ctx *gin.Context) {
	longURL := ctx.PostForm("long_url")

	if longURL == "" {
		utils.RenderTemplate(200, ctx, components.ShortenerResult("", "Long URL is required"))
		return
	}

	if len(longURL) > 2000 { // 2000 is the maximum length of a URL
		utils.RenderTemplate(200, ctx, components.ShortenerResult("", "Long URL is too long"))
		return
	}

	shortUrlCode := utils.GetRandomCode(5)
	baseUrl := env.GetEnv().ApplicationBaseUrl
	shortURL := baseUrl + "/r/" + shortUrlCode

	fmt.Println(longURL, shortURL)

	db, err := db.GetDb()

	if err != nil {
		utils.RenderTemplate(200, ctx, components.ShortenerResult("", "Please try again later"))
		return
	}
	_, err = db.Exec(`INSERT INTO short_urls (long_url, short_url) VALUES ($1, $2)`, longURL, shortUrlCode)

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	utils.RenderTemplate(200, ctx, components.ShortenerResult(shortURL, ""))
}

func (sc ShortenerController) Redirect(c *gin.Context) {
	shortUrlCode := c.Param("short_url_code")
	db, err := db.GetDb()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	var longUrl string
	err = db.QueryRow(`SELECT long_url FROM short_urls WHERE short_url = $1`, shortUrlCode).Scan(&longUrl)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.Redirect(302, longUrl)
}
