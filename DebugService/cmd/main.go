package main

import (
	"os"
	"flag"
	"time"
	"DebugService/dist"
	"fmt"
	"syscall"
	"context"
	"os/signal"
	"DebugService/internal/service"
	"DebugService/internal/server/grpc"
	"DebugService/internal/server/http"
	"github.com/mapgoo-lab/atreus/pkg/log"
	"github.com/mapgoo-lab/atreus/pkg/conf/paladin"
)

func main() {
	flag.Parse()
	if err := paladin.Init(); err != nil {
		panic(err)
	}
	log.Init(nil) // debug flag: log.dir={path}
	defer log.Close()
	log.Info(fmt.Sprintf("DebugService start, version:%s, build date:%s",dist.Version,dist.Build_date))
	svc := service.New()
	grpcSrv := grpc.New(svc)
	httpSrv := http.New(svc)
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Info("DebugService get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			ctx, cancel := context.WithTimeout(context.Background(), 35*time.Second)
			if err := grpcSrv.Shutdown(ctx); err != nil {
				log.Error("DebugService grpcSrv.Shutdown error(%v)", err)
			}
			if err := httpSrv.Shutdown(ctx); err != nil {
				log.Error("DebugService httpSrv.Shutdown error(%v)", err)
			}
			log.Info(fmt.Sprintf("DebugService exit, version:%s, build date:%s",dist.Version,dist.Build_date))
			svc.Close()
			cancel()
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
