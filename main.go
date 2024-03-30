package main

import (
	"fmt"
	"log"
	"syscall"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"

	"github.com/raighneweng/tinyurl-go/pkg/gredis"
	"github.com/raighneweng/tinyurl-go/pkg/setting"
	"github.com/raighneweng/tinyurl-go/routers"
)

func init() {
	setting.Setup()
	gredis.Setup()
}

func main() {
	gin.SetMode(setting.ServerSetting.RunMode)

	routersInit := routers.InitRouter()

	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	maxHeaderBytes := 1 << 20

	log.Printf("listen and serve on 0.0.0.0:%v", setting.ServerSetting.HttpPort)

	// If you want Graceful Restart, you need a Unix system and download github.com/fvbock/endless
	endless.DefaultReadTimeOut = readTimeout
	endless.DefaultWriteTimeOut = writeTimeout
	endless.DefaultMaxHeaderBytes = maxHeaderBytes

	server := endless.NewServer(fmt.Sprintf(":%v", setting.ServerSetting.HttpPort), routersInit)

	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}
}
