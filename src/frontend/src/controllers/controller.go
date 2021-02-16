package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandlerHomePage(ctx *gin.Context) {
	ctx.HTML(
		http.StatusOK,
		"index.html",
		gin.H{
			"message": "Welcome! REWARDING SITES!",
		},
	)
}
