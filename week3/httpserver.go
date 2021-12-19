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

	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

func main() {
	g, ctx := errgroup.WithContext(context.Background())
	mux := http.NewServeMux()
	// 简单的http请求处理，返回 Hello, world!
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, world!")
	})
	// 同步channel，接收关闭服务的请求
	closeChan := make(chan bool)
	// 关闭http服务
	mux.HandleFunc("/close", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Good bye~")
		closeChan <- true
	})
	// 创建一个http服务
	server := http.Server{
		Handler: mux,
		Addr:    ":8080",
	}
	// 启动http服务
	g.Go(func() error {
		return server.ListenAndServe()
	})
	// 关闭http服务
	g.Go(func() error {
		select {
		case <-ctx.Done():
			log.Println("errgroup exit")
		case <-closeChan:
			log.Println("close server")
		}
		timeoutCtx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		log.Println("closing server")
		return server.Shutdown(timeoutCtx)
	})
	// 捕获 os 退出信号
	g.Go(func() error {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		select {
		case <-ctx.Done():
			return ctx.Err()
		case sig := <-quit:
			return errors.Errorf("Get signal: %v", sig)
		}
	})
	fmt.Printf("errgroup exiting: %+v\n", g.Wait())
}
