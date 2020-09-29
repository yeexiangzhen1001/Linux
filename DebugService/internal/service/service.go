package service

import (
	"DebugService/internal/dao"
	"context"
	"github.com/mapgoo-lab/atreus/pkg/conf/paladin"
	"github.com/mapgoo-lab/atreus/pkg/net/trace/zipkin"
)

// Service service.
type Service struct {
	ac  *paladin.Map
	dao *dao.Dao
}

type traceConfig struct {
	Zipkin *zipkin.Config
}

// New new a service and return.
func New() (s *Service) {
	var ac = new(paladin.TOML)
	if err := paladin.Watch("application.toml", ac); err != nil {
		panic(err)
	}

	var kin traceConfig
	if err := paladin.Get("application.toml").UnmarshalTOML(&kin); err == nil {
		zipkin.Init(kin.Zipkin)
	}

	s = &Service{
		ac:  ac,
		dao: dao.New(),
	}
	return s
}

// Ping ping the resource.
func (s *Service) Ping(ctx context.Context) (err error) {
	return s.dao.Ping(ctx)
}

// Close close the resource.
func (s *Service) Close() {
	s.dao.Close()
}
