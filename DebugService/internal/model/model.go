package model

import (
	pb "DebugService/api"
	"strings"
)

const (
	Success         pb.PAAS_ERROR_CODE = pb.PAAS_ERROR_CODE_COMM_SUCCESS
	Failed          pb.PAAS_ERROR_CODE = pb.PAAS_ERROR_CODE_COMM_ECODE_OTHER
	ConnectFailed   pb.PAAS_ERROR_CODE = pb.PAAS_ERROR_CODE_COMM_ECODE_CONNECT_ERROR
	NoData          pb.PAAS_ERROR_CODE = pb.PAAS_ERROR_CODE_COMM_ECODE_NO_DATA
	SqlQueryError   pb.PAAS_ERROR_CODE = pb.PAAS_ERROR_CODE_DAP_ECODE_SQL_QUERY_ERROR
	SqlExecError    pb.PAAS_ERROR_CODE = pb.PAAS_ERROR_CODE_DAP_ECODE_SQL_EXEC_ERROR
	ScanRowError    pb.PAAS_ERROR_CODE = pb.PAAS_ERROR_CODE_DAP_ECODE_ROW_SCAN_ERROR
	InvalidParam    pb.PAAS_ERROR_CODE = pb.PAAS_ERROR_CODE_COMM_ECODE_INVALID_PARAM_ERROR
	SqlUnknownError pb.PAAS_ERROR_CODE = pb.PAAS_ERROR_CODE_DAP_ECODE_SQL_UNKNOWN_ERROR
)

func Len(str string) int {
	str = strings.TrimSpace(str)
	return len(str)
}