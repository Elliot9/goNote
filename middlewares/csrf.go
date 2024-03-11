package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/csrf"
	adapter "github.com/gwatts/gin-adapter"
)

func CSRF() gin.HandlerFunc {
	return adapter.Wrap(csrf.Protect([]byte("32-byte-long-auth-key")))
}

func CsrfToken() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Header("X-CSRF-Token", csrf.Token(context.Request))
	}
}
