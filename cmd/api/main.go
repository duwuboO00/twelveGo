package main

import (
	"net/http"
	"twelveGo/config"
	"twelveGo/internal/database"
	"twelveGo/pkg/api"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// GetJWTClaims 為一個 stub，請根據實際情況實作
func GetJWTClaims(c *gin.Context) interface{} {
	// TODO: 完整實作 JWT 解析
	return map[string]interface{}{}
}

func main() {
	cfg := config.GetConfig()
	router := gin.Default()

	// 設置基於 Cookie 的會話
	store := cookie.NewStore([]byte("your_secret_key"))
	router.Use(sessions.Sessions("twelveGo_session", store))

	// 設置 HTML 模板
	router.LoadHTMLGlob("web/templates/*")

	// 使用 database.InitDB() 建立 DB 連線，內部使用的連線字串
	// 已在 database.DBURL 中組裝好，作為 singleton 使用
	db, err := database.InitDB()
	if err != nil {
		zap.L().Fatal("無法初始化資料庫", zap.Error(err))
	}
	// 為避免 db 未使用錯誤，暫時忽略 db 變數，未來如需使用，請將 db 傳遞至相關模組
	_ = db
	zap.L().Info("資料庫初始化成功！")

	// 取得從 database.DBURL 產生的 dbURL
	dbURL := database.DBURL

	// 使用 config 中 MongoDB 設定建立 URI
	mongoURL := cfg.MongoDBConfig.URI

	// 路由設置
	router.GET("/", func(c *gin.Context) {
		session := sessions.Default(c)
		userID := session.Get("user_id")
		jwtClaims := GetJWTClaims(c)

		c.HTML(http.StatusOK, "index.html", gin.H{
			"Port":      cfg.Port,
			"DBURL":     dbURL, // 顯示由 database.InitDB() 組裝好的 dbURL
			"MongoURL":  mongoURL,
			"UserID":    userID,
			"JWTClaims": jwtClaims,
		})
	})

	router.GET("/heartbeat", api.Heartbeat)
	router.POST("/login", api.Login)
	router.GET("/memberheartbeat", api.MemberHeartbeat)

	// 自定義 404 處理
	router.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "404.html", nil)
	})

	// 全局錯誤處理中介層
	router.Use(func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			c.HTML(http.StatusInternalServerError, "error.html",
				gin.H{"errors": c.Errors})
		}
	})

	router.Run(":" + cfg.Port)
}
