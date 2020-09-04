package main

import (
	"fmt"
	"github.com/marshhu/ma-novel-crawler/router"
	"golang.org/x/net/context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	Run()
}

func Run() {
	router := router.InitRouter()

	host := fmt.Sprintf(":%d", 8086)
	s := &http.Server{
		Addr:           host,
		Handler:        router,
		ReadTimeout:    time.Duration(60) * time.Second,
		WriteTimeout:   time.Duration(60) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	go func() {
		log.Println("Server Listen at:" + host)
		if err := s.ListenAndServe(); err != nil {
			log.Printf("Listen:%s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Shutdown Server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
