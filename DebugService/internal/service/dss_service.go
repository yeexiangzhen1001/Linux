package service

import (
	pb "DebugService/api"
	"context"
	"github.com/mapgoo-lab/atreus/pkg/log"
)

//获取设备路由信息
func (s *Service) GetRouterInfo(ctx context.Context, req *pb.GetRouterInfoReq) (*pb.GetRouterInfoResp, error) {
	log.Error("GetRouterInfo(req:%v).", req)
	resp, _ := s.dao.GetRouterInfo(ctx, req)
	return resp, nil
}
