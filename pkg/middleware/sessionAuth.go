package middleware

import (
	"net/http"
	"twelveGo/pkg/models"

	"github.com/gin-gonic/gin"
)

// SessionAuth 驗證 session cookie
func SessionAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		sid, err := c.Cookie("session_id")
		if err != nil || sid == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "未授權"})
			return
		}

		s, err := models.GetSession(sid)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "會話查詢失敗"})
			return
		}
		if s == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "無效的會話"})
			return
		}

		_ = models.UpdateSessionActivity(sid)
		c.Set("user_id", s.UserID)
		c.Next()
	}
}
