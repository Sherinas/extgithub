package main

import (
	handler "login.go/handler"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {

	route := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	route.Use(sessions.Sessions("mysession", store))

	route.Static("/static", "./static")
	route.LoadHTMLGlob("template/**")

	route.GET("/", handler.GetLogin)
	route.POST("/", handler.GetData)
	route.GET("/dash", handler.GetHome)
	route.GET("/logout", handler.LogOut)

	route.Run(":5000")

	// main barch edit

}
