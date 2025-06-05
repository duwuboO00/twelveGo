package api

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// Heartbeat 用於健康檢查
func Heartbeat(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}

// googleAuth 處理 Google 認證邏輯
func googleAuth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Google Auth"})
}

// Login 處理用戶登入請求
func Login(c *gin.Context) {
	type loginRequest struct {
		Identifier string `json:"identifier" form:"identifier" binding:"required"`
		Password   string `json:"password" form:"password" binding:"required"`
	}
	var req loginRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "參數錯誤"})
		return
	}

	// TODO: 從資料庫取出使用者資料並驗證密碼
	if req.Identifier != "demo" || req.Password != "password" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "帳號或密碼錯誤"})
		return
	}

	session := sessions.Default(c)
	session.Set("user_id", req.Identifier)
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "無法建立會話"})
		return
	}

	fakeUser := gin.H{
		"id":    req.Identifier,
		"email": "demo@example.com",
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "登入成功",
		"user":    fakeUser,
	})
}

// Logout 處理用戶登出請求
func Logout(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get("user_id")
	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "無效的會話",
		})
		return
	}
	session.Delete("user_id")
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "會話無法刪除",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "登出成功",
	})
}

// InitAuthRouter 初始化認證路由
func InitAuthRouter(r *gin.Engine) {
	auth := r.Group("/auth")
	{
		auth.GET("/google", googleAuth)
		auth.POST("/login", Login)
		auth.POST("/logout", Logout)
	}

	// 添加 Heartbeat 路由
	r.GET("/heartbeat", Heartbeat)
}
