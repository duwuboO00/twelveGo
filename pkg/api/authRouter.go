package api

// import (
// 	"config" // 引入配置包
// 	"github.com/gin-contrib/sessions"
// 	"github.com/gin-gonic/gin"
// 	"internal/database" // 假設這是數據庫操作的包
// 	"net/http"
// 	"pkg/models" // 引入 User 模型
// )

// Heartbeat 用於健康檢查
// func Heartbeat(c *gin.Context) {
// 	c.String(http.StatusOK, "ok")
// }

// // Login 處理用戶登入請求
// func Login(c *gin.Context) {
// 	// 讀取配置信息
// 	cfg, err := config.LoadConfig()
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "無法加載配置"})
// 		return
// 	}

// 	// 在這裡實現用戶登入的邏輯
// 	// 例如，驗證用戶憑證，創建會話等

// 	c.JSON(http.StatusOK, gin.H{"message": "登入成功"})
// }

// // Logout 處理用戶登出請求
// func Logout(c *gin.Context) {
// 	session := sessions.Default(c)
// 	user := session.Get("user_id")
// 	if user == nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "無效的會話"})
// 		return
// 	}
// 	session.Delete("user_id")
// 	if err := session.Save(); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "會話無法刪除"})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"message": "登出成功"})
// }
