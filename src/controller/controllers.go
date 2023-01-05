package controller

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Set("id", 111)
	session.Set("email", "test@mail.ru")
	session.Save()
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Use Sign In successfully",
	})

}

func Logout(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Clear()
	session.Save()
	ctx.JSON(http.StatusOK, gin.H{
		"message": "User Sign out successfully",
	})
}
