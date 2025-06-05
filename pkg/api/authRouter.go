package api

import (
	"net/http"
	"twelveGo/config"

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
	cfg, err := config.LoadConfig()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "無法加載配置",
		})
		return
	}

	// 解析 Google ID
	var payload struct {
		GoogleID string `form:"google_id" json:"google_id" binding:"required"`
	}
	if err := c.ShouldBind(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "缺少 Google ID"})
		return
	}

	// TODO: 應檢查資料庫或其他認證機制確認 google_id 是否有效
	if payload.GoogleID != "google_abcdef" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "無效的 Google ID"})
		return
	}
	session := sessions.Default(c)
	session.Set("user_id", payload.GoogleID)
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "無法建立會話"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "登入成功",
		// 注意：正式版本請移除下方配置回應
		"config": cfg,
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
