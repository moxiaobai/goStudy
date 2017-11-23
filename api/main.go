package main

import (
	"context"
	"github.com/moxiaobai/goStudy/api/routers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	db "github.com/moxiaobai/goStudy/api/database"
	 "github.com/moxiaobai/goStudy/api/config"
)



func main() {
	defer db.SqlDB.Close()

	//初始化路由
	router := routers.InitRoutes()

	//获取配置
	config := config.GetConfig()

	srv := &http.Server{
		Addr:   config.Host,
		Handler: router,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
