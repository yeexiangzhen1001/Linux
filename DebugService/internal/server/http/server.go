package http

import (
	"DebugService/api"
	"DebugService/internal/service"
	"DebugService/internal/utility"
	"context"
	"fmt"
	"github.com/mapgoo-lab/atreus/pkg/conf/env"
	"github.com/mapgoo-lab/atreus/pkg/conf/paladin"
	"github.com/mapgoo-lab/atreus/pkg/log"
	"github.com/mapgoo-lab/atreus/pkg/naming"
	"github.com/mapgoo-lab/atreus/pkg/naming/etcd"
	bm "github.com/mapgoo-lab/atreus/pkg/net/http/blademaster"
	"net/http"
	"os"
	"time"
)

type Server struct {
	engine  *bm.Engine
	builder *etcd.EtcdBuilder
}

var (
	svc *service.Service
)

// New new a bm server.
func New(s *service.Service) *Server {
	var (
		hc struct {
			Server *bm.ServerConfig
		}
	)
	if err := paladin.Get("application.toml").UnmarshalTOML(&hc); err != nil {
		if err != paladin.ErrNotExist {
			panic(err)
		}
	}
	svc = s
	engine := bm.DefaultServer(hc.Server)
	api.RegisterDebugServiceBMServer(engine, s)

	initRouter(engine)
	if err := engine.Start(); err != nil {
		panic(err)
	}
	appid := fmt.Sprintf("http-%s", env.AppID)
	zone := env.Zone
	hostname := fmt.Sprintf("http-%s-%s-%d-%d", env.AppID, env.Hostname, time.Now().Unix(), os.Getpid())
	addrs := utility.GetArrary("HTTP_ADDR", "http://192.168.110.11:8107", "http.addr", "etcd http addr", ",")

	builder, err := etcd.New(nil)
	if err == nil {
		//注册服务发现
		builder.Build(appid)
		_, err := builder.Register(context.Background(), &naming.Instance{
			AppID:    appid,
			Hostname: hostname,
			Zone:     zone,
			Addrs:    addrs,
		})

		if err != nil {
			log.Error("DebugService start error:%s", err.Error())
			panic(err)
		}
	}

	log.Error("DebugService init http etcd server(appid:%s,hostname:%s).", appid, hostname)

	return &Server{
		engine:  engine,
		builder: builder,
	}
}

func (s *Server) Shutdown(ctx context.Context) (err error) {
	s.builder.Close()
	return s.engine.Shutdown(ctx)
}

func initRouter(e *bm.Engine) {
	e.Ping(ping)
	g := e.Group("/DebugService")
	{
		g.GET("/start", HelloWorld)
	}
}

func ping(ctx *bm.Context) {
	if err := svc.Ping(ctx); err != nil {
		log.Error("ping error(%v)", err)
		ctx.AbortWithStatus(http.StatusServiceUnavailable)
	}
}

// example for http request handler.
func HelloWorld(c *bm.Context) {
	c.JSON("app-api is running...", nil)
}
