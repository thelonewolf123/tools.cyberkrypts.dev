package utils

import (
	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
)

func RenderTemplate(status int, ctx *gin.Context, component templ.Component) {
	ctx.Status(status)
	component.Render(ctx, ctx.Writer)
}
