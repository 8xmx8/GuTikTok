package main

import (
	"GuTikTok/config"
	"GuTikTok/logging"
	"GuTikTok/utils/tracing"
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	// Set Trace Provider
	tp, err := tracing.SetTraceProvider(config.WebServiceName)

	if err != nil {
		logging.Logger.WithFields(log.Fields{
			"err": err,
		}).Panicf("Error to set the trace")
	}
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			logging.Logger.WithFields(log.Fields{
				"err": err,
			}).Errorf("Error to set the trace")
		}
	}()

	g := gin.Default()
	//url

	g.GET("ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Run Server
	RunServer(g)

}
func RunServer(g *gin.Engine) {
	base := config.Conf.Server.Address + config.WebServiceAddr
	log.Infof("启动服务器 @ %s", base)
	srv := &http.Server{Addr: base, Handler: g}
	go func() {
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
