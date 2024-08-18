package handler

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"login.go/cache"
)

func GetHome(ctx *gin.Context) {
	cache.ClearCache(ctx)
	session := sessions.Default(ctx)

	userID := session.Get("userID")
	if userID == nil {
		ctx.Redirect(http.StatusSeeOther, "/")
		return
	}

	ctx.HTML(http.StatusOK, "home.html", nil)
}
