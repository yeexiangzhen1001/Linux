#!/bin/bash

#代码根目录
src_root=`pwd`

env_install() {
    yum install -y epel-release
    yum install -y snap
    snap install --classic protobuf
    go get -u github.com/mapgoo-lab/atreus/tool/protobuf/protoc-gen-bm
    go get -u github.com/gogo/protobuf/protoc-gen-gofast
    go get -u github.com/mapgoo-lab/atreus/tool/protobuf/protoc-gen-bswagger
    go get -u github.com/go-swagger/go-swagger/cmd/swagger
}

build_interface() {

    cd $src_root/api
    pb_path='../../interface'

    if [ $# == 2 ]; then
        pb_path=$1
    fi

    rm -rf mapgoo.paas.*

    #编译grpc接口文件
    protoc --proto_path=$GOPATH/src \
           --proto_path=$GOPATH/src/github.com/mapgoo-lab/atreus/third_party \
           --proto_path=$src_root/api \
           --proto_path=$pb_path \
           --gofast_out=plugins=grpc:. \
           $pb_path/mapgoo.paas.ecode.proto \
           $pb_path/mapgoo.paas.base.proto \
           $pb_path/mapgoo.paas.baseex.proto \
           $pb_path/mapgoo.paas.status.data.proto \
           $pb_path/mapgoo.paas.travel.data.proto \
           $pb_path/mapgoo.paas.live.data.proto \
           $pb_path/mapgoo.paas.video.data.proto \
           $pb_path/mapgoo.paas.debug.api.proto \
           $pb_path/mapgoo.paas.debug.api.data.proto \
           $pb_path/mapgoo.paas.dss.proto \
           $pb_path/mapgoo.paas.device.proto \
           $pb_path/mapgoo.paas.dap.device.proto \
           $pb_path/mapgoo.paas.cap.device.proto \
           $pb_path/mapgoo.paas.dap.status.proto \
           $pb_path/mapgoo.paas.cap.status.proto \
           $pb_path/mapgoo.paas.socol.data.proto \
           $pb_path/mapgoo.paas.cap.socol.proto \
           $pb_path/mapgoo.paas.dms.proto \
           $pb_path/mapgoo.paas.hold.data.proto \
           $pb_path/mapgoo.paas.hms.proto \
           $pb_path/mapgoo.paas.travel.data.proto \
           $pb_path/mapgoo.paas.live.data.proto \
           $pb_path/mapgoo.paas.vehicle.data.proto \
           $pb_path/mapgoo.paas.vms.proto \
           $pb_path/mapgoo.paas.ecode.proto \
           $pb_path/mapgoo.paas.cds.proto \
           $pb_path/mapgoo.paas.cap.hold.proto \
           $pb_path/mapgoo.paas.dap.hold.proto

    #编译http接口文件
    protoc --proto_path=$GOPATH/src \
           --proto_path=$GOPATH/src/github.com/mapgoo-lab/atreus/third_party \
           --proto_path=$src_root/api \
           --proto_path=$pb_path \
           --bm_out=explicit_http=true:.  \
           $pb_path/mapgoo.paas.debug.api.proto

     if [ -f "./mapgoo.paas.debug.api.bm.go" ]; then
        sed -i "s/mapgoo_paas_debug_api_data.//g" mapgoo.paas.debug.api.bm.go
     fi
}

build() {
    build_interface
    cd $src_root
    if [ ! -d "$src_root/dist" ]; then
        mkdir -p $src_root/dist
    fi

    if [ -f "$src_root/dist/version.go" ]; then
        rm -rf $src_root/dist/version.go
    fi

    touch $src_root/dist/version.go

    DATE=$(date +"%Y-%m-%d %H:%M:%S")
    echo "package dist" > $src_root/dist/version.go
    echo "" >> $src_root/dist/version.go
    echo "var (" >> $src_root/dist/version.go
    echo "  Version string = \"$service_version\"" >> $src_root/dist/version.go
    echo "  Build_date string = \"$DATE\"" >> $src_root/dist/version.go
    echo ")" >> $src_root/dist/version.go

    go build -o debugService cmd/main.go
}

run() {
    cd $src_root

    export APP_ID=DebugService
    export LOG_V=1
    export LOG_STDOUT=true
    export LOG_DIR=/data/paas/logs/DebugService
    export HTTP_PERF=tcp://0.0.0.0:29003
    export GRPC="tcp://0.0.0.0:9103/?timeout=60s&idle_timeout=60s"
    export HTTP="tcp://0.0.0.0:8103/?timeout=60s"
    export ETCD_ENDPOINTS=192.168.110.31:2379
    export ETCD_PREFIX=mapgoo-pass
#    export DAP_TARGET=direct://default/127.0.0.1:9100
#    export CAP_TARGET=direct://default/127.0.0.1:9101
     export DMS_TARGET=direct://default/192.168.110.211:9103
#    export UMS_TARGET=direct://default/127.0.0.1:9106
     export DSS_TARGET=direct://default/192.168.110.211:9104
#    export HMS_TARGET=direct://default/127.0.0.1:9105
#    export VMS_TARGET=direct://default/127.0.0.1:9111
#    export CMS_TARGET=direct://default/127.0.0.1:9102

    export DAP_TARGET=etcd://default/dap-go
    export CAP_TARGET=etcd://default/cap-go
#   export DMS_TARGET=etcd://default/dms
    export UMS_TARGET=etcd://default/ums
#   export DSS_TARGET=etcd://default/dss
    export HMS_TARGET=etcd://default/hms
    export VMS_TARGET=etcd://default/vms
    export CMS_TARGET=etcd://default/cms

    export GRPC_ADDR=grpc://192.168.110.31:9103
    export HTTP_ADDR=http://192.168.110.31:8103
    export ICEPROXY_TARGET=direct://default/192.168.110.31:11003

    ./debugService -conf configs-dev/
}


print_help() {
    echo "Usage: "
    echo "  $0 env ---Install env"
    echo "  $0 version version_str ---Build all as a version"
    echo "  $0 interface --- Build interface file with default pb path"
    echo "  $0 interface /path/pb/interface --- Build interface file with specified pb path"
    echo "  $0 run --- Run the service"
}

case $1 in
    env)
        echo "install env..."
        env_install
        ;;
    version)
        if [ $# -lt 2 ]; then
            echo $#
            print_help
            exit
        fi

        if [ $# == 2 ]; then
            build $2
        fi
        ;;
    interface)
        if [ $# == 2 ]; then
            echo "build interface with pb path ..."
            build_interface $2
        else
            echo "build interface ..."
            build_interface
        fi
        ;;
    run)
        run
        ;;
    *)

    print_help
        ;;
esac
