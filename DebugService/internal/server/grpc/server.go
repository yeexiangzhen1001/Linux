package grpc

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
	"github.com/mapgoo-lab/atreus/pkg/net/rpc/warden"
	"os"
	"time"
)

type Server struct {
	ws      *warden.Server
	builder *etcd.EtcdBuilder
}

// New new a grpc server.
func New(svc *service.Service) *Server {
	var rc struct {
		Server *warden.ServerConfig
	}
	if err := paladin.Get("application.toml").UnmarshalTOML(&rc); err != nil {
		if err != paladin.ErrNotExist {
			panic(err)
		}
	}
	ws := warden.NewServer(rc.Server)
	api.RegisterDebugServiceServer(ws.Server(), svc)
	ws, err := ws.Start()
	if err != nil {
		panic(err)
	}

	appid := env.AppID
	zone := env.Zone
	hostname := fmt.Sprintf("grpc-%s-%s-%d-%d", appid, env.Hostname, time.Now().Unix(), os.Getpid())
	addrs := utility.GetArrary("GRPC_ADDR", "grpc://192.168.110.31:9107", "grpc.addr", "etcd grpc addr", ",")

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
			panic(err)
		}
	}
	log.Error(fmt.Sprintf("DebugService init grpc etcd server(appid:%s,hostname:%s).", appid, hostname))

	return &Server{
		ws:      ws,
		builder: builder,
	}
}

func (s *Server) Shutdown(ctx context.Context) (err error) {
	s.builder.Close()
	return s.ws.Shutdown(ctx)
}
