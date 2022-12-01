// Package main 服务主要入口
package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	Documents "transformer_mz/api/documents"
	"transformer_mz/libs/cron"
	"transformer_mz/libs/message"
	"transformer_mz/libs/third"
	"transformer_mz/middleware/cors"

	"transformer_mz/databases/connector"
	"transformer_mz/libs/config"
	"transformer_mz/libs/log"
	"transformer_mz/libs/permission"
	"transformer_mz/routers"
)

type ServerConfigWrap struct {
	IP   string `yaml:"ip"`
	Port string `yaml:"port"`
}

type ServerConfig struct {
	ServerConfigWrap `yaml:"server"`
}

func main() {
	serverConfig := &ServerConfig{}
	err := config.Load(serverConfig)
	if err != nil {
		log.Fatalln(err)
	}
	err = log.InitLog()
	if err != nil {
		log.Fatalln(err)
	}
	err = connector.InitMysqlConnector()
	if err != nil {
		log.Fatalln(err)
	}
	err = message.Init()
	if err != nil {
		log.Fatalln(err)
	}
	err = Documents.InitDocuments()
	if err != nil {
		log.Fatalln(err)
	}
	err = permission.Init()
	if err != nil {
		log.Fatalln(err)
	}
	err = third.InitThirdParty()
	if err != nil {
		log.Fatalln(err)
	}
	cron.InitCron()

	server := &http.Server{
		Addr:    serverConfig.IP + ":" + serverConfig.Port,
		Handler: cors.New(routers.SetupRouter()),
	}
	log.Print("服务启动中，服务地址:", server.Addr)
	go func() {
		err = server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Print(err)
		}
	}()
	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM)
	select {
	case sig := <-s:
		signal.Reset(syscall.SIGINT, syscall.SIGTERM)
		log.Printf("捕获到信号%v，准备停止服务, 再按一次\"ctrl+c\"可强制退出", sig)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		err = server.Shutdown(ctx)
		if err != nil {
			log.Print("服务退出失败:", err)
		}
		cancel()
		db, err := connector.DataSource.DB()
		if err != nil {
			log.Print(err)
		} else {
			err = db.Close()
			if err != nil {
				log.Print(err)
			}
		}
		log.Print("服务退出完毕")
		os.Exit(1)
	}
}
