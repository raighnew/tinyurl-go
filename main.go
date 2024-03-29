package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/raighneweng/tinyurl-go/controller"
	"github.com/raighneweng/tinyurl-go/pkg/gredis"
	"github.com/raighneweng/tinyurl-go/pkg/setting"
)

func init() {
	setting.Setup()
	gredis.Setup()
}

func main() {
	gin.SetMode(setting.ServerSetting.RunMode)
	handler := gin.Default()

	handler.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	handler.POST("/generate", controller.Generate)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%v", setting.ServerSetting.HttpPort),
		Handler: handler,
	}

	go func() {
		// service connections
		log.Printf("listen and serve on 0.0.0.0:%v", setting.ServerSetting.HttpPort)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	defaultTimeout := time.Second

	if setting.ServerSetting.Environment == "Production" {
		defaultTimeout = 10 * time.Second
	}

	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 10 seconds.
	select {
	case <-ctx.Done():
		log.Println("timeout of 10 seconds.")
	}

	log.Println("Server exiting")
}
