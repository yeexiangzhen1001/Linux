package dao

import (
	pb "DebugService/api"
	"DebugService/internal/model"
	"context"
	"flag"
	"github.com/mapgoo-lab/atreus/pkg/conf/paladin"
	"github.com/mapgoo-lab/atreus/pkg/log"
	"github.com/mapgoo-lab/atreus/pkg/naming/etcd"
	"github.com/mapgoo-lab/atreus/pkg/net/rpc/warden"
	"github.com/mapgoo-lab/atreus/pkg/net/rpc/warden/resolver"
	"google.golang.org/grpc"
	"os"
)

// Dao dao interface
type Dao struct {
	DssConn pb.DSSClient
	DmsConn pb.DMSClient
}

func checkErr(err error) {
	if err != nil {
		log.Error("Dss.Dao.Error:%s", err.Error())
		panic(err)
	}
}

var (
		_dssTarget string
		_dmsTarget string
)

const (
	DSS_TARGET         = "DSS_TARGET"
	DSS_TARGET_DEFAULT = "direct://default/192.168.110.72:10001"

	DMS_TARGET         = "DMS_TARGET"
	DMS_TARGET_DEFAULT = "direct://default/192.168.110.72:10002"

)

func addFlag(fs *flag.FlagSet) {


	fs.StringVar(&_dssTarget, "dss.target", getEnvOrDefault(DSS_TARGET, DSS_TARGET_DEFAULT), "dss target address or env variable.")
	fs.StringVar(&_dmsTarget, "dms.target", getEnvOrDefault(DMS_TARGET, DMS_TARGET_DEFAULT), "dms target address or env variable.")
}

func getEnvOrDefault(envName string, val string) string {
	envValue := os.Getenv(envName)
	if model.Len(envValue) > 0 {
		return envValue
	}
	return val
}

func getLocalConn(c *warden.ClientConfig, target string, opts ...grpc.DialOption) *grpc.ClientConn {
	client := warden.NewClient(c, opts...)
	conn, err := client.Dial(context.Background(), target)

	if nil != err {
		log.Error(target+".getLocalConn.Error:%s", err.Error())
		panic(err)
	}
	return conn
}

func DSSClient(c *warden.ClientConfig, opts ...grpc.DialOption) pb.DSSClient {
	conn := getLocalConn(c, _dssTarget, opts...)
	return pb.NewDSSClient(conn)
}
func DMSClient(c *warden.ClientConfig, opts ...grpc.DialOption) (pb.DMSClient) {
	conn := getLocalConn(c, _dmsTarget, opts...)
	return pb.NewDMSClient(conn)
}
// New new a dao and return.
func New() *Dao {
	var df struct {
		Dss *warden.ClientConfig
		Dms *warden.ClientConfig
	}

	checkErr(paladin.Get("application.toml").UnmarshalTOML(&df))
	checkErr(paladin.Get("application.toml").UnmarshalTOML(&model.GlobalConfig))

	addFlag(flag.CommandLine)

	//客户端注册服务发现
	e, err := etcd.New(nil)
	if err != nil {
		panic(err)
	}
	resolver.Register(e)

	daohandle := new(Dao)
	daohandle.DssConn = DSSClient(df.Dss)
	daohandle.DmsConn = DMSClient(df.Dms)

	return daohandle
}

// Close close the resource.
func (d *Dao) Close() {
}

// Ping ping the resource.
func (d *Dao) Ping(ctx context.Context) (err error) {
	return
}
