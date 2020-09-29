package dao

import (
	"context"
	pb "DebugService/api"
	"github.com/mapgoo-lab/atreus/pkg/ecode"
	"github.com/prometheus/common/log"
)

//获取路由信息
func (d *Dao) GetRouterInfo(ctx context.Context, rep *pb.GetRouterInfoReq) (*pb.GetRouterInfoResp,error) {
	req1 := new(pb.GetDeviceRouterReq)
	req1.DeviceId = rep.DeviceId

	resp, errDss := d.DssConn.GetDeviceRouter(ctx, req1)
	if errDss != nil {
		log.Error("GetDeviceRouter from DSS failed, error:%+v", errDss)
		return nil, errDss
	}

	if resp.Base.Error != pb.PAAS_ERROR_CODE_COMM_SUCCESS {
		log.Error("GetDeviceRouter from DSS failed, error code:%d, reasion:%s", resp.Base.Error, resp.Base.Reason)
		return nil, ecode.New(int(pb.PAAS_ERROR_CODE_COMM_SUCCESS))
	}

	out := &pb.GetRouterInfoResp{
		RouterInfo: resp.Router,
	}

	return out,nil
}
