package main

import (
	"context"
	"echo-playground/conf"
	ctl "echo-playground/controller"
	rt "echo-playground/router"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

var (
	g errgroup.Group
)
var configFlag = flag.String("config", "./conf/config.toml", "toml file to use for configuration")

func main() {
	flag.Parse()
	cf := conf.NewConfig(*configFlag)

	// TODO: Add logger
	log.Default().Println("Server starting...")

	// TODO: Add model
	if controller, err := ctl.NewCTL(cf); err != nil {
		panic(fmt.Errorf("Failed to create controller: %v", err))
	} else if router, err := rt.NewRouter(cf, controller); err != nil {
		panic(fmt.Errorf("Failed to create router: %v", err))
	} else {
		mapi := &http.Server{
			Addr:           cf.Server.Port,
			Handler:        router.Idx(),
			ReadTimeout:    5 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		}

		g.Go(func() error {
			return mapi.ListenAndServe()
		})

		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		log.Default().Println("Server shutting down...")

		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		if err := mapi.Shutdown(ctx); err != nil {
			log.Default().Fatalf("Server forced to shutdown: %v", err)
		}

		select {
		case <-ctx.Done():
			log.Default().Println("timeout of 2 seconds.")
		}
		log.Default().Println("Server exiting")
	}
	if err := g.Wait(); err != nil {
		log.Default().Fatalf("Server error: %v", err)
	}
}
