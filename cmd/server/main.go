package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"hex_ddd_cqs_example/config"
	http2 "hex_ddd_cqs_example/http"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	conf, err := config.New()
	if err != nil {
		log.Fatal(err)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(gin.Recovery())

	r.GET("/", http2.Homepage())

	server := &http.Server{
		Addr:    fmt.Sprintf("%s", conf.ApiHost()),
		Handler: r,
	}
	defer func() { _ = server.Shutdown(ctx) }()

	fmt.Println("system start")
	defer func() {
		fmt.Println("system stop")
	}()

	readyCh := make(chan struct{})
	go func() {
		fmt.Println("system ready to serve")
		<-readyCh
	}()

	/* signal handling */
	go func() {
		done := make(chan os.Signal, 1)
		signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		<-done
		fmt.Println("signal trapped")
		cancel()
	}()

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}
