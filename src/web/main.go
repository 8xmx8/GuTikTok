package main

import (
	"GuTikTok/config"
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	g := gin.Default()

	//url

	base := fmt.Sprintf("%s:%d", config.Conf.Server.Address, config.Conf.Server.Port) //127.0.0.1:23927
	log.Infof("启动服务器 @ %s", base)
	srv := &http.Server{Addr: base, Handler: g}

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		var err error
		if config.Conf.Server.Https {
			err = srv.ListenAndServeTLS(config.Conf.Server.CertFile, config.Conf.Server.KeyFile)
		} else {
			err = srv.ListenAndServe()
		}
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("无法启动: %s", err.Error())
		}
	}()

	wg.Wait()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
