package api

import (
	"net/http"
	"twelveGo/config"
	"twelveGo/pkg/middleware"
	"twelveGo/pkg/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "無法加載配置"})
		return
	}

	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "參數錯誤"})
		return
	}

	userID := uuid.New()
	sessionID := uuid.New().String()
	if err := models.CreateSession(userID, sessionID, c.ClientIP(), c.GetHeader("User-Agent")); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "建立 session 失敗"})
		return
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "session_id",
		Value:    sessionID,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	})

	c.JSON(http.StatusOK, gin.H{"message": "登入成功", "config": cfg})
}

// LoginTest 生成測試用 session 並回傳 ok
func LoginTest(c *gin.Context) {
	userID := uuid.New()
	sessionID := uuid.New().String()
	if err := models.CreateSession(userID, sessionID, c.ClientIP(), c.GetHeader("User-Agent")); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "建立 session 失敗"})
		return
	}
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "session_id",
		Value:    sessionID,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	})
	c.String(http.StatusOK, "ok")
}

// Logout 處理用戶登出請求
func Logout(c *gin.Context) {
	sid, err := c.Cookie("session_id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "無效的會話"})
		return
	}
	if err := models.DeleteSession(sid); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "會話無法刪除"})
		return
	}
	http.SetCookie(c.Writer, &http.Cookie{Name: "session_id", MaxAge: -1, Path: "/"})
	c.JSON(http.StatusOK, gin.H{"message": "登出成功"})
}

// CheckSession 測試登入狀態
func CheckSession(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}

// InitAuthRouter 初始化認證路由
func InitAuthRouter(r *gin.Engine) {
	auth := r.Group("/auth")
	{
		auth.GET("/google", googleAuth)
		auth.POST("/login", Login)
		auth.GET("/login-test", LoginTest)
		auth.POST("/logout", Logout)
		auth.GET("/check", middleware.SessionAuth(), CheckSession)
	}

	// 添加 Heartbeat 路由
	r.GET("/heartbeat", Heartbeat)
}
