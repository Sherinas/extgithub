package handler

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"login.go/cache"
)

func LogOut(ctx *gin.Context) {
	cache.ClearCache(ctx)
	session := sessions.Default(ctx)
	session.Clear()
	session.Save()
	ctx.Redirect(http.StatusSeeOther, "/")

}
