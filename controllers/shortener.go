package controllers

import (
	"fmt"
	"math/rand"
	"youtube-downloader-web/db"
	"youtube-downloader-web/env"

	"github.com/gin-gonic/gin"
)

type ShortenerController struct{}

func generateShortUrlCode() string {
	shortUrlLength := 5
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	shortUrl := ""
	for i := 0; i < shortUrlLength; i++ {
		shortUrl += string(chars[rand.Intn(len(chars))])
	}
	return shortUrl
}

func (sc ShortenerController) Index(ctx *gin.Context) {
	ctx.HTML(200, "shortener.html", gin.H{})
}

func (sc ShortenerController) Generate(ctx *gin.Context) {
	longURL := ctx.PostForm("long_url")
	shortUrlCode := generateShortUrlCode()
	baseUrl := env.GetEnv().ApplicationBaseUrl
	shortURL := baseUrl + "/r/" + shortUrlCode

	fmt.Println(longURL, shortURL)

	db, err := db.GetDb()

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	_, err = db.Exec(`INSERT INTO short_urls (long_url, short_url) VALUES ($1, $2)`, longURL, shortUrlCode)

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.HTML(200, "short-url.html", gin.H{
		"short_url": shortURL,
	})
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
