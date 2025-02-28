package api

import (
	"net/http"
	"twelveGo/config" // 引入配置包
	// "twelveGo/internal/database" // 數據庫操作包
	// "twelveGo/pkg/models"        // 引入模型包

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// Heartbeat 用於健康檢查
func Heartbeat(c *gin.Context) {
	c.String(http.StatusOK, "ok")
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

	// 在這裡實現用戶登入邏輯，例如驗證用戶憑證和創建會話
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
