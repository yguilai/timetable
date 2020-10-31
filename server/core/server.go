package core

import (
	"ccsu/global"
	"ccsu/initialize"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func RunServer() {
	if global.Config.System.Mode != "dev" {
		gin.SetMode(gin.ReleaseMode)
	}

	if global.Config.Redis.Enable {
		initialize.Redis()
	}

	Router := initialize.Routers()

	s := initServer(fmt.Sprintf(":%s", global.Config.System.Port), Router)
	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			global.Log.Error(err.Error())
		}
	}()

	// 利用http.Server的shutdown优雅的重启
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		global.Log.Error(fmt.Sprintf("Server shutdown: %s", err.Error()))
	}
	global.Log.Info("Server exiting")
}

func initServer(addr string, router *gin.Engine) *http.Server {
	return &http.Server{
		Addr:           addr,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}