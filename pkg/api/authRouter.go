package api

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Heartbeat(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}

// Login 處理用戶登入請求
func Login(c *gin.Context) {
	// 從請求中獲取用戶名和密碼
	var loginInfo struct {
		Username string `form:"username" json:"username" binding:"required"`
		Password string `form:"password" json:"password" binding:"required"`
	}
	if err := c.ShouldBind(&loginInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 模擬用戶認證，實際應用中應替換為您的用戶認證邏輯
	if loginInfo.Username != "admin" || loginInfo.Password != "password" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "認證失敗"})
		return
	}

	// 認證成功，設置會話
	session := sessions.Default(c)
	session.Set("user_id", loginInfo.Username)
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "無法保存會話"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "登入成功"})
}

// Logout 處理用戶登出請求
func Logout(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get("user_id")
	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "無效的會話"})
		return
	}
	session.Delete("user_id")
	session.Save()
	c.JSON(http.StatusOK, gin.H{"message": "登出成功"})
}

// SomeProtectedRoute 是一個示例受保護的路由處理函數
func SomeProtectedRoute(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get("user_id")
	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未授權"})
		return
	}

	// 執行受保護的操作
	c.JSON(http.StatusOK, gin.H{"message": "這是一個受保護的內容"})
}
