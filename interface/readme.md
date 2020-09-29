# 接口定义规范

## RPC接口定义规范

### RPC接口定义方式
1. 因为使用的微服务框架是grpc，所以使用protobuf接口定义方式
2. 因为grpc的限定，接口定义必须使用protobuf3，文档参考：[Language Guide (proto3)](https://developers.google.com/protocol-buffers/docs/proto3)

### RPC接口定义规范

1. pb文件命名方式:mapgoo.paas.${modula_name}.proto
2. pb文件起始一行，显式加上使用pb3语法
```
syntax = "proto3";
```
2. 基本数据类型，能在各个模块通用的数据，比如卡信息、卡用量、卡套餐信息、设备信息、用户信息，定义到maogoo.paas.base.proto文件中
3. 模块服务的接口定义，放在mapgoo.paas.${service_name}.proto
4. 模块专用的数据类型，特别是模块接口入参和出参，如果特别多，可以单独定义到mapgoo.paas.${service_name}.data.proto文件中，如果不是很多，可以定义到mapgoo.paas.${service_name}.proto
5. 根据go语言特性，每个pb定义文件中，需要加上
```go
option go_package = "api";
``` 
用来指定编译成golang之后的包名
6. pb的包名定义方式：mapgoo.paas.${modula_name|servce_name}
7. 类型命名规则：大写字母开头，使用驼峰命名法
8. 字段命名规则：全部小写字母，使用下划线_连接
9. 接口命名规则：接口名使用大写字母开头，使用驼峰命名法
10. 接口定义规则：原则上一个接口对应一个request类型和一个response类型
11. 为了方便调试和生成文档，接口定义必须添加生成http接口的option选项，例子：
```
//设备添加
rpc addDeviceInfo(SaveDeviceInfoReq) returns (SaveDeviceInfoResp) {
    option (google.api.http) = {
        post:"/dms/addDeviceInfo"
    };
}
```
7. pb文件的注释，因为依赖swagger生成文档，请在被注释的对象上一行使用'//注释内容的方式'，并在注释的对象空格一行，以下为例子：
```
//设备ID
message DeviceId
{
	//设备逻辑id
    uint32 object_id = 1;

    //设备imei
    string imei = 2;
}
```


### RPC接口定义范例

* 基本数据类型定义，在maogoo.paas.base.proto文件中定义，以设备信息为例

```
syntax = "proto3";                //第一行，明确指定使用pb3的语法

package mapgoo.paas.dms;        //第二行，明确指定pb文件的包名

option go_package = "api";     //指定生成golang代码的包名

option (gogoproto.goproto_getters_all) = false;

import "github.com/gogo/protobuf/gogoproto/gogo.proto";
import "google/api/annotations.proto";

//设备ID
message DeviceId
{
	//设备逻辑id
    uint32 object_id = 1;

    //设备imei
    string imei = 2;
}

//设备类型
message DeviceTypeInfo 
{
	//设备类型ID
	uint32 mdt_type_id = 1;	

	//是否为无线设备		
	uint32 is_wireless = 2;	

	//设备品牌	
	string brand = 3;	

	//设备协议			
	string protocol = 4;

	//设备厂家				
	string factory = 5;	

	//是否为智能设备			
	uint32 is_smart = 6;		

	//备注		
	string remark = 7;				
}

message GetDeviceTypeReq 
{
	//设备ID
	DeviceId device_id = 1;
}

message GetDeviceTypeResp
{
	//基础响应
	mapgoo.paas.base.BaseRespInfo base = 1 [(gogoproto.jsontag) = 'base'];

	//设备类型信息
	DeviceTypeInfo type_info = 2 [(gogoproto.jsontag) = 'type_info'];
}

service DMS {

	//添加设备
	 rpc getDeviceType(GetDeviceTypeReq) returns (GetDeviceTypeResp) {
        option (google.api.http) = {
            post:"/dms/addDeviceInfo"
        };
    }
}

```




