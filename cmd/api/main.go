package main

import (
	"net/http"
	"twelveGo/config"
	"twelveGo/internal/database"
	"twelveGo/pkg/api"
	"twelveGo/pkg/models"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	cfg := config.GetConfig()
	router := gin.Default()

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
		sid, _ := c.Cookie("session_id")
		var sessionData string
		if sid != "" {
			if s, err := models.GetSession(sid); err == nil && s != nil {
				sessionData = s.UserID.String()
			}
		}
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Port":        cfg.Port,
			"DBURL":       dbURL,
			"MongoURL":    mongoURL,
			"SessionID":   sid,
			"SessionData": sessionData,
		})
	})

	// 初始化路由
	api.InitAuthRouter(router)
	api.InitShoppingCartRouter(router)
	api.InitMemberRouter(router)

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
