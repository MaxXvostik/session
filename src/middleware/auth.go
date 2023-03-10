package middleware

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		sessionId := session.Get("id")
		if sessionId == nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": "unauthorized",
			})
			ctx.Abort()
		}
	}
}
