package middleware

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func SessionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := c.Request.Cookie("ice_session_id")
		if errors.Is(err, http.ErrNoCookie) {
			c.SetCookie("ice_session_id", time.Now().String(), 3600, "/", "localhost", false, true)
		}
		c.Next()
	}
}
