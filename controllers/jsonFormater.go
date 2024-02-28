package controllers

import (
	"github.com/gin-gonic/gin"
	"tools.cyberkrypts.dev/templates/pages"
	"tools.cyberkrypts.dev/utils"
)

type JsonFormatterController struct{}

func (c JsonFormatterController) Index(ctx *gin.Context) {
	utils.RenderTemplate(200, ctx, pages.JsonFormatterIndex())
}
