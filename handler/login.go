package handler

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"login.go/cache"
)

func GetLogin(ctx *gin.Context) {

	cache.ClearCache(ctx)
	session := sessions.Default(ctx)
	userID := session.Get("userID")

	if userID != nil {

		ctx.Redirect(http.StatusSeeOther, "/dash")
		return
	}

	ctx.HTML(http.StatusOK, "index.html", gin.H{})
}

func GetData(ctx *gin.Context) {
	cache.ClearCache(ctx)
	session := sessions.Default(ctx)

	username := ctx.PostForm("user")
	password := ctx.PostForm("pswd")

	if username == "user" && password == "12345" {

		session.Set("userID", username)
		session.Save()

		ctx.Redirect(http.StatusSeeOther, "/dash")

		return
	}
	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"error": "Invalid username or password",
	})

	ctx.Redirect(http.StatusSeeOther, "/")
	///a

}
