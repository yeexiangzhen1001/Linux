#!/bin/bash

#代码根目录
src_root=`pwd`


build_interface() {

    #删除所有编译好的文档
    rm -rf api
    mkdir api

    #编译app-api接口文档
    protoc --proto_path=$GOPATH/src \
           --proto_path=$GOPATH/src/github.com/mapgoo-lab/atreus/third_party \
           --proto_path=. --bswagger_out=explicit_http=true:./api \
           mapgoo.paas.app.api.proto \
           mapgoo.paas.device.api.proto \
           mapgoo.paas.dms.proto \
           mapgoo.paas.dss.proto \
           mapgoo.paas.hms.proto \
           mapgoo.paas.liveservice.proto \
           mapgoo.paas.lss.proto \
           mapgoo.paas.mfs.proto \
           mapgoo.paas.openplatform.api.proto \
           mapgoo.paas.packagemanager.proto \
           mapgoo.paas.rars.proto \
           mapgoo.paas.reportservice.proto \
           mapgoo.paas.cms.simscache.proto \
           mapgoo.paas.commonservice.proto \
           mapgoo.paas.socol.data.gateway.proto \
           mapgoo.paas.socol.task.manager.proto \
           mapgoo.paas.ssvp.proto \
           mapgoo.paas.cds.proto \
           mapgoo.paas.ums.proto \
           mapgoo.paas.vms.proto \
           mapgoo.paas.dap.alarm.proto \
           mapgoo.paas.dap.device.proto \
           mapgoo.paas.dap.hold.proto \
           mapgoo.paas.dap.live.proto \
           mapgoo.paas.dap.packagemanager.proto \
           mapgoo.paas.dap.sims.proto \
           mapgoo.paas.dap.status.proto \
           mapgoo.paas.dap.travel.proto \
           mapgoo.paas.dap.user.proto \
           mapgoo.paas.dap.vehicle.proto \
           mapgoo.paas.cap.device.proto \
           mapgoo.paas.cap.hold.proto \
           mapgoo.paas.cap.package.proto \
           mapgoo.paas.cap.sims.proto \
           mapgoo.paas.cap.socol.proto \
           mapgoo.paas.cap.status.proto \
           mapgoo.paas.cap.tool.proto \
           mapgoo.paas.cap.travel.proto \
           mapgoo.paas.cap.user.proto \
           mapgoo.paas.socol.estimate.api.proto
}

build_interface
