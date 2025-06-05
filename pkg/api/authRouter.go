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

	// 解析使用者提交的帳號與密碼
	var creds struct {
		Username string `form:"username" json:"username" binding:"required"`
		Password string `form:"password" json:"password" binding:"required"`
	}
	if err := c.ShouldBind(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "缺少帳號或密碼"})
		return
	}

	// TODO: 這裡應改為檢查資料庫中的帳號密碼
	if creds.Username != "user123" || creds.Password != "password123" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "帳號或密碼錯誤"})
		return
	}

	// 建立 session
	session := sessions.Default(c)
	session.Set("user_id", creds.Username)
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
