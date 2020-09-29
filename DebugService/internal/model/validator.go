package model

import (
	pb "DebugService/api"
	"DebugService/internal/utility"
)

var DBState bool = false

func SetBaseRespInfo(error int, args ...interface{}) (resp *pb.BaseRespInfo) {
	isEmpty := len(args) == 0
	args = append(args, "")
	var reason2 = ""
	if error == 0 {
		reason2 = utility.GetFlagValue(isEmpty, "操作成功!", args[0]).(string)
	} else if error > 0 {
		reason2 = utility.GetFlagValue(isEmpty, "操作失败!", args[0]).(string)
	} else {
		reason2 = utility.GetFlagValue(isEmpty, "操作异常!", args[0]).(string)
	}
	resp = &pb.BaseRespInfo{Error: pb.PAAS_ERROR_CODE(error), Reason: reason2}
	return
}

func CheckBaseRespInfo(req interface{}) (resp *pb.BaseRespInfo) {
	if nil == req {
		resp = &pb.BaseRespInfo{Error: 1, Reason: "参数为空!"}
	} else {
		resp = &pb.BaseRespInfo{Error: 0, Reason: "参数正常"}
	}

	return
}

func ResetBaseRespInfo(info *pb.BaseRespInfo, err pb.PAAS_ERROR_CODE, msg string) {
	if nil != info {
		info.Error = (err)
		info.Reason = msg
	}
}
