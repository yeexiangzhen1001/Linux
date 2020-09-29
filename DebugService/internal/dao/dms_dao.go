package dao

import (
	"github.com/mapgoo-lab/atreus/pkg/log"
	pb "DebugService/api"
	"context"
	"github.com/mapgoo-lab/atreus/pkg/ecode"
)

//获取设备基本信息
/*func (d *Dao) GetDeviceBaseInfo(ctx context.Context, objectid uint32, imei string) (*pb.DeviceBaseInfoResp) {
	req := new(pb.DeviceInfoReq)
	req.DeviceId = new(pb.DeviceId)
	req.DeviceId.ObjectId = objectid
	req.DeviceId.Imei = imei
	resp, err := d.DmsConn.GetDeviceBaseInfo(ctx, req)
	if err != nil {
		log.Error("GetDeviceBaseInfo failed(objectid:%d,imei:%s,err:%v).", objectid, imei, err)
		resp = new(pb.DeviceBaseInfoResp)
		resp.Base = new(pb.BaseRespInfo)
		resp.Base.Error = model.Failed
		resp.Base.Reason = "获取设备信息异常"
		return resp
	} else if resp.Base.Error != model.Success {
		log.Error("GetDeviceBaseInfo failed(objectid:%d,imei:%s,resp:%v).", objectid, imei, resp)
		if resp.BaseInfo == nil {
			resp.Base.Error = model.NoData
			resp.Base.Reason = "设备信息不存在"
		}

		return resp
	}

	return resp
}*/
//获取设备基本信息
func (d *Dao)  GetDeviceInfo(ctx context.Context, rep *pb.GetDeviceInfoReq) (*pb.GetDeviceBaseInfoResp,error) {
	req1 := new(pb.DeviceInfoReq)
	req1.DeviceId = rep.DeviceId

	resp, errDms := d.DmsConn.GetDeviceBaseInfo(ctx, req1)
	if errDms != nil {
		log.Error("GetDeviceBaseInfo from DMS failed, error:%+v", errDms)
		return nil, errDms
	}

	if resp.Base.Error != pb.PAAS_ERROR_CODE_COMM_SUCCESS {
		log.Error("GetDeviceBaseInfo from DMS failed, error code:%d, reasion:%s", resp.Base.Error, resp.Base.Reason)
		return nil, ecode.New(int(pb.PAAS_ERROR_CODE_COMM_SUCCESS))
	}

	out := &pb.GetDeviceBaseInfoResp{
		BaseInfo: resp.BaseInfo,
	}

	return out,nil
}