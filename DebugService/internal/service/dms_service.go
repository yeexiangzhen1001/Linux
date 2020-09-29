package service
import (
	pb "DebugService/api"
	"context"
	"github.com/mapgoo-lab/atreus/pkg/log"

)

/*func (s *Service) GetDeviceBaseInfo(ctx context.Context, req *pb.GetDeviceInfoReq) (resp *pb.GetDeviceBaseInfoResp) {
	//log.Error("GetDeviceBaseInfo(req:%v).", req)
	objectid := req.DeviceId.ObjectId
	imei := req.DeviceId.Imei
	resp1 := s.dao.GetDeviceBaseInfo(ctx, objectid, imei)
	resp.BaseInfo = resp1.BaseInfo

	return resp
	//resp = new(pb.GetDeviceBaseInfoResp)
	//	resp = s.dao.GetDeviceBaseInfo(ctx, req.DeviceId)
	//log.Info("baseinfo:", resp.BaseInfo, " respbaseinfo:", resp.Base)
}*/
//获取设备基本信息
func (s *Service) GetDeviceInfo(ctx context.Context, req *pb.GetDeviceInfoReq) (*pb.GetDeviceBaseInfoResp, error) {
	//log.Error("GetDeviceBaseInfo(req:%v).", req)
	log.Error("GetDeviceInfo(req:%v).", req)
	resp, _ := s.dao.GetDeviceInfo(ctx, req)
	return resp, nil

}
