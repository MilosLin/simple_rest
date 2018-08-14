package api

import (
	"log"
	"net/http"

	"simple_rest/config"
	"simple_rest/env"

	"github.com/gin-gonic/gin"
)

// Start RESTful API Service
func Start() {
	router := gin.New()

	// 使用Recovery中間件，避免panic導致伺服器中斷運行
	router.Use(gin.Recovery())

	// 綁定路由與控制方法
	BindRouting(router)

	// 啟動 HTTP 伺服器
	server := &http.Server{
		Addr:    config.Forge().GetString(env.ApiListenPort),
		Handler: router,
	}

	log.Fatalf("Server exiting. %v", server.ListenAndServe())
}
