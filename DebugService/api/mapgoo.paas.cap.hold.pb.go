// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mapgoo.paas.cap.hold.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("mapgoo.paas.cap.hold.proto", fileDescriptor_d79c7237bb6ec1ae) }

var fileDescriptor_d79c7237bb6ec1ae = []byte{
	// 424 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x95, 0xcd, 0x4a, 0xeb, 0x40,
	0x14, 0xc7, 0x1b, 0x2e, 0xdc, 0xc5, 0xc0, 0xa5, 0xdc, 0xa1, 0x55, 0xa9, 0x90, 0x7d, 0xab, 0xa4,
	0x5f, 0x4f, 0x60, 0x5b, 0x88, 0xa2, 0x8b, 0xe2, 0xa8, 0x0b, 0x77, 0x63, 0x33, 0x6d, 0x82, 0x35,
	0x73, 0xcc, 0x0c, 0x6a, 0x05, 0x41, 0xf0, 0x25, 0x7c, 0x02, 0x9f, 0xc5, 0xa5, 0x8f, 0x20, 0xf5,
	0x45, 0x24, 0x1f, 0xd8, 0x36, 0x9d, 0xa4, 0x19, 0xb7, 0xc9, 0xef, 0x9c, 0xdf, 0xc9, 0x99, 0xe1,
	0x1f, 0x54, 0xbb, 0xa1, 0x30, 0xe1, 0xdc, 0x02, 0x4a, 0x85, 0x35, 0xa2, 0x60, 0xb9, 0x7c, 0xea,
	0x58, 0x10, 0x70, 0xc9, 0x71, 0x45, 0xf5, 0xae, 0xb6, 0xbb, 0xfc, 0x34, 0xa2, 0x1d, 0x2a, 0x69,
	0x5c, 0xd2, 0x79, 0xfb, 0x87, 0xf0, 0x21, 0x9f, 0x3a, 0x47, 0xfe, 0x98, 0xf7, 0xe9, 0xc8, 0x65,
	0xc3, 0x80, 0x3f, 0xcc, 0xb0, 0x8f, 0xfe, 0xdb, 0x4c, 0x1e, 0xc0, 0x75, 0x9f, 0xfb, 0x63, 0x6f,
	0x12, 0xbd, 0xc1, 0x7b, 0x96, 0xba, 0xd3, 0x1a, 0x79, 0xca, 0x6e, 0x6b, 0xfb, 0xc5, 0x61, 0x01,
	0xf8, 0x1e, 0x55, 0xce, 0xc1, 0xa1, 0x92, 0xa5, 0x94, 0x56, 0x46, 0x17, 0x15, 0x1c, 0x5a, 0x9b,
	0x5a, 0xbc, 0x00, 0xec, 0xa2, 0xf2, 0x80, 0x4d, 0x99, 0x64, 0xd1, 0xa3, 0x70, 0x0b, 0xb8, 0x9e,
	0xd1, 0x23, 0xc5, 0x85, 0xba, 0x46, 0x51, 0x54, 0x00, 0x7e, 0x31, 0xd0, 0xce, 0xda, 0xc7, 0x5f,
	0xb0, 0x40, 0x78, 0xdc, 0xc7, 0x9d, 0xa2, 0xdb, 0x4a, 0x0a, 0x42, 0x79, 0x57, 0xbb, 0x26, 0x99,
	0x82, 0xe8, 0x4e, 0x41, 0x7e, 0x31, 0x05, 0xc9, 0x9b, 0xc2, 0x45, 0x65, 0xc2, 0xe4, 0x8a, 0xbb,
	0x9e, 0xdd, 0x27, 0xad, 0x6c, 0x14, 0x45, 0x63, 0x93, 0x5d, 0xd0, 0x64, 0x17, 0x37, 0xd9, 0x0a,
	0xd3, 0x13, 0xda, 0x22, 0xf4, 0x6e, 0x71, 0xc7, 0x4e, 0x3c, 0x11, 0x43, 0xb8, 0x95, 0x35, 0xaf,
	0x12, 0x0f, 0xbd, 0x6d, 0xcd, 0x0a, 0x01, 0xf8, 0x11, 0x55, 0x97, 0x0f, 0x7e, 0x61, 0x6f, 0x16,
	0xb8, 0x26, 0x2b, 0xf2, 0x96, 0x5e, 0x81, 0x80, 0x9f, 0xb4, 0x80, 0x63, 0x36, 0xeb, 0xcd, 0xa2,
	0x38, 0x19, 0xe4, 0xa7, 0xc5, 0x32, 0xb9, 0x31, 0x2d, 0x56, 0x61, 0x01, 0xf8, 0xd9, 0x40, 0xdb,
	0xc9, 0x2a, 0x08, 0x93, 0xd2, 0xf3, 0x27, 0x03, 0x6f, 0x94, 0x7c, 0xee, 0x86, 0xd5, 0xa5, 0xf9,
	0x50, 0xde, 0xd1, 0x2d, 0x89, 0x4f, 0x3b, 0xde, 0xc7, 0xda, 0x00, 0xf9, 0xeb, 0x53, 0xf9, 0xdb,
	0x9a, 0x15, 0x71, 0x5e, 0xda, 0x4c, 0x86, 0x2b, 0x39, 0x0b, 0x18, 0x1b, 0x52, 0xe9, 0xe6, 0xe7,
	0xa5, 0x0a, 0xce, 0xcb, 0x4b, 0x35, 0x1f, 0x8b, 0x89, 0x8e, 0x98, 0x68, 0x8a, 0x49, 0x86, 0xb8,
	0x57, 0x7d, 0x9f, 0x9b, 0xc6, 0xc7, 0xdc, 0x34, 0x3e, 0xe7, 0xa6, 0xf1, 0xfa, 0x65, 0x96, 0x2e,
	0xff, 0x50, 0xf0, 0x86, 0xa5, 0xab, 0xbf, 0xd1, 0x8f, 0xac, 0xfb, 0x1d, 0x00, 0x00, 0xff, 0xff,
	0x1b, 0x33, 0x7d, 0x92, 0x19, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HoldInfoCacheProxyClient is the client API for HoldInfoCacheProxy service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HoldInfoCacheProxyClient interface {
	//获取渠道配置
	GetApkConfigCache(ctx context.Context, in *GetApkConfigCacheReq, opts ...grpc.CallOption) (*GetApkConfigCacheResp, error)
	//修改渠道配置
	UpdateApkConfigCache(ctx context.Context, in *UpdateApkConfigCacheReq, opts ...grpc.CallOption) (*UpdateApkConfigCacheResp, error)
	//删除缓存信息
	DeleteCacheInfo(ctx context.Context, in *DeleteCacheInfoReq, opts ...grpc.CallOption) (*DeleteCacheInfoResp, error)
	//获取渠道配置版本号
	GetApkConfigCacheVersion(ctx context.Context, in *GetApkConfigCacheVersionReq, opts ...grpc.CallOption) (*GetApkConfigCacheVersionResp, error)
	//设置渠道配置版本号
	SetApkConfigCacheVersion(ctx context.Context, in *SetApkConfigCacheVersionReq, opts ...grpc.CallOption) (*SetApkConfigCacheVersionResp, error)
	//公共设置缓存
	SetCacheVersion(ctx context.Context, in *SetCacheVersionReq, opts ...grpc.CallOption) (*SetCacheVersionResp, error)
	//公共获取缓存
	GetCacheVersion(ctx context.Context, in *GetCacheVersionReq, opts ...grpc.CallOption) (*GetCacheVersionResp, error)
	//设置渠道配置缓存
	SaveApkConfigListCache(ctx context.Context, in *SaveApkConfigListCacheReq, opts ...grpc.CallOption) (*SaveApkConfigListCacheResp, error)
	//获取渠道配置缓存
	GetApkConfigListCache(ctx context.Context, in *GetApkConfigListCacheReq, opts ...grpc.CallOption) (*GetApkConfigListCacheResp, error)
	//获取渠道key
	GetAppKeyByHoldID(ctx context.Context, in *GetAppKeyByHoldIDReq, opts ...grpc.CallOption) (*GetAppKeyByHoldIDResp, error)
	//缓存渠道配置字典
	SaveApkSettingDictCache(ctx context.Context, in *SaveApkSettingDictCacheReq, opts ...grpc.CallOption) (*SaveApkSettingDictCacheResp, error)
	//获取渠道配置字典缓存
	GetApkSettingDictCache(ctx context.Context, in *GetApkSettingDictCacheReq, opts ...grpc.CallOption) (*GetApkSettingDictCacheResp, error)
	//获取渠道目录树缓存
	GetHoldTreePathCache(ctx context.Context, in *GetHoldTreePathCacheReq, opts ...grpc.CallOption) (*GetHoldTreePathCacheResp, error)
	//保存渠道目录树缓存
	SetHoldTreePathCache(ctx context.Context, in *SetHoldTreePathCacheReq, opts ...grpc.CallOption) (*SetHoldTreePathCacheResp, error)
}

type holdInfoCacheProxyClient struct {
	cc *grpc.ClientConn
}

func NewHoldInfoCacheProxyClient(cc *grpc.ClientConn) HoldInfoCacheProxyClient {
	return &holdInfoCacheProxyClient{cc}
}

func (c *holdInfoCacheProxyClient) GetApkConfigCache(ctx context.Context, in *GetApkConfigCacheReq, opts ...grpc.CallOption) (*GetApkConfigCacheResp, error) {
	out := new(GetApkConfigCacheResp)
	err := c.cc.Invoke(ctx, "/mapgoo.paas.cap.hold.HoldInfoCacheProxy/GetApkConfigCache", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *holdInfoCacheProxyClient) UpdateApkConfigCache(ctx context.Context, in *UpdateApkConfigCacheReq, opts ...grpc.CallOption) (*UpdateApkConfigCacheResp, error) {
	out := new(UpdateApkConfigCacheResp)
	err := c.cc.Invoke(ctx, "/mapgoo.paas.cap.hold.HoldInfoCacheProxy/UpdateApkConfigCache", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *holdInfoCacheProxyClient) DeleteCacheInfo(ctx context.Context, in *DeleteCacheInfoReq, opts ...grpc.CallOption) (*DeleteCacheInfoResp, error) {
	out := new(DeleteCacheInfoResp)
	err := c.cc.Invoke(ctx, "/mapgoo.paas.cap.hold.HoldInfoCacheProxy/DeleteCacheInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *holdInfoCacheProxyClient) GetApkConfigCacheVersion(ctx context.Context, in *GetApkConfigCacheVersionReq, opts ...grpc.CallOption) (*GetApkConfigCacheVersionResp, error) {
	out := new(GetApkConfigCacheVersionResp)
	err := c.cc.Invoke(ctx, "/mapgoo.paas.cap.hold.HoldInfoCacheProxy/GetApkConfigCacheVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *holdInfoCacheProxyClient) SetApkConfigCacheVersion(ctx context.Context, in *SetApkConfigCacheVersionReq, opts ...grpc.CallOption) (*SetApkConfigCacheVersionResp, error) {
	out := new(SetApkConfigCacheVersionResp)
	err := c.cc.Invoke(ctx, "/mapgoo.paas.cap.hold.HoldInfoCacheProxy/SetApkConfigCacheVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *holdInfoCacheProxyClient) SetCacheVersion(ctx context.Context, in *SetCacheVersionReq, opts ...grpc.CallOption) (*SetCacheVersionResp, error) {
	out := new(SetCacheVersionResp)
	err := c.cc.Invoke(ctx, "/mapgoo.paas.cap.hold.HoldInfoCacheProxy/SetCacheVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *holdInfoCacheProxyClient) GetCacheVersion(ctx context.Context, in *GetCacheVersionReq, opts ...grpc.CallOption) (*GetCacheVersionResp, error) {
	out := new(GetCacheVersionResp)
	err := c.cc.Invoke(ctx, "/mapgoo.paas.cap.hold.HoldInfoCacheProxy/GetCacheVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *holdInfoCacheProxyClient) SaveApkConfigListCache(ctx context.Context, in *SaveApkConfigListCacheReq, opts ...grpc.CallOption) (*SaveApkConfigListCacheResp, error) {
	out := new(SaveApkConfigListCacheResp)
	err := c.cc.Invoke(ctx, "/mapgoo.paas.cap.hold.HoldInfoCacheProxy/SaveApkConfigListCache", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *holdInfoCacheProxyClient) GetApkConfigListCache(ctx context.Context, in *GetApkConfigListCacheReq, opts ...grpc.CallOption) (*GetApkConfigListCacheResp, error) {
	out := new(GetApkConfigListCacheResp)
	err := c.cc.Invoke(ctx, "/mapgoo.paas.cap.hold.HoldInfoCacheProxy/GetApkConfigListCache", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *holdInfoCacheProxyClient) GetAppKeyByHoldID(ctx context.Context, in *GetAppKeyByHoldIDReq, opts ...grpc.CallOption) (*GetAppKeyByHoldIDResp, error) {
	out := new(GetAppKeyByHoldIDResp)
	err := c.cc.Invoke(ctx, "/mapgoo.paas.cap.hold.HoldInfoCacheProxy/GetAppKeyByHoldID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *holdInfoCacheProxyClient) SaveApkSettingDictCache(ctx context.Context, in *SaveApkSettingDictCacheReq, opts ...grpc.CallOption) (*SaveApkSettingDictCacheResp, error) {
	out := new(SaveApkSettingDictCacheResp)
	err := c.cc.Invoke(ctx, "/mapgoo.paas.cap.hold.HoldInfoCacheProxy/SaveApkSettingDictCache", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *holdInfoCacheProxyClient) GetApkSettingDictCache(ctx context.Context, in *GetApkSettingDictCacheReq, opts ...grpc.CallOption) (*GetApkSettingDictCacheResp, error) {
	out := new(GetApkSettingDictCacheResp)
	err := c.cc.Invoke(ctx, "/mapgoo.paas.cap.hold.HoldInfoCacheProxy/GetApkSettingDictCache", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *holdInfoCacheProxyClient) GetHoldTreePathCache(ctx context.Context, in *GetHoldTreePathCacheReq, opts ...grpc.CallOption) (*GetHoldTreePathCacheResp, error) {
	out := new(GetHoldTreePathCacheResp)
	err := c.cc.Invoke(ctx, "/mapgoo.paas.cap.hold.HoldInfoCacheProxy/GetHoldTreePathCache", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *holdInfoCacheProxyClient) SetHoldTreePathCache(ctx context.Context, in *SetHoldTreePathCacheReq, opts ...grpc.CallOption) (*SetHoldTreePathCacheResp, error) {
	out := new(SetHoldTreePathCacheResp)
	err := c.cc.Invoke(ctx, "/mapgoo.paas.cap.hold.HoldInfoCacheProxy/SetHoldTreePathCache", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HoldInfoCacheProxyServer is the server API for HoldInfoCacheProxy service.
type HoldInfoCacheProxyServer interface {
	//获取渠道配置
	GetApkConfigCache(context.Context, *GetApkConfigCacheReq) (*GetApkConfigCacheResp, error)
	//修改渠道配置
	UpdateApkConfigCache(context.Context, *UpdateApkConfigCacheReq) (*UpdateApkConfigCacheResp, error)
	//删除缓存信息
	DeleteCacheInfo(context.Context, *DeleteCacheInfoReq) (*DeleteCacheInfoResp, error)
	//获取渠道配置版本号
	GetApkConfigCacheVersion(context.Context, *GetApkConfigCacheVersionReq) (*GetApkConfigCacheVersionResp, error)
	//设置渠道配置版本号
	SetApkConfigCacheVersion(context.Context, *SetApkConfigCacheVersionReq) (*SetApkConfigCacheVersionResp, error)
	//公共设置缓存
	SetCacheVersion(context.Context, *SetCacheVersionReq) (*SetCacheVersionResp, error)
	//公共获取缓存
	GetCacheVersion(context.Context, *GetCacheVersionReq) (*GetCacheVersionResp, error)
	//设置渠道配置缓存
	SaveApkConfigListCache(context.Context, *SaveApkConfigListCacheReq) (*SaveApkConfigListCacheResp, error)
	//获取渠道配置缓存
	GetApkConfigListCache(context.Context, *GetApkConfigListCacheReq) (*GetApkConfigListCacheResp, error)
	//获取渠道key
	GetAppKeyByHoldID(context.Context, *GetAppKeyByHoldIDReq) (*GetAppKeyByHoldIDResp, error)
	//缓存渠道配置字典
	SaveApkSettingDictCache(context.Context, *SaveApkSettingDictCacheReq) (*SaveApkSettingDictCacheResp, error)
	//获取渠道配置字典缓存
	GetApkSettingDictCache(context.Context, *GetApkSettingDictCacheReq) (*GetApkSettingDictCacheResp, error)
	//获取渠道目录树缓存
	GetHoldTreePathCache(context.Context, *GetHoldTreePathCacheReq) (*GetHoldTreePathCacheResp, error)
	//保存渠道目录树缓存
	SetHoldTreePathCache(context.Context, *SetHoldTreePathCacheReq) (*SetHoldTreePathCacheResp, error)
}

// UnimplementedHoldInfoCacheProxyServer can be embedded to have forward compatible implementations.
type UnimplementedHoldInfoCacheProxyServer struct {
}

func (*UnimplementedHoldInfoCacheProxyServer) GetApkConfigCache(ctx context.Context, req *GetApkConfigCacheReq) (*GetApkConfigCacheResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetApkConfigCache not implemented")
}
func (*UnimplementedHoldInfoCacheProxyServer) UpdateApkConfigCache(ctx context.Context, req *UpdateApkConfigCacheReq) (*UpdateApkConfigCacheResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateApkConfigCache not implemented")
}
func (*UnimplementedHoldInfoCacheProxyServer) DeleteCacheInfo(ctx context.Context, req *DeleteCacheInfoReq) (*DeleteCacheInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCacheInfo not implemented")
}
func (*UnimplementedHoldInfoCacheProxyServer) GetApkConfigCacheVersion(ctx context.Context, req *GetApkConfigCacheVersionReq) (*GetApkConfigCacheVersionResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetApkConfigCacheVersion not implemented")
}
func (*UnimplementedHoldInfoCacheProxyServer) SetApkConfigCacheVersion(ctx context.Context, req *SetApkConfigCacheVersionReq) (*SetApkConfigCacheVersionResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetApkConfigCacheVersion not implemented")
}
func (*UnimplementedHoldInfoCacheProxyServer) SetCacheVersion(ctx context.Context, req *SetCacheVersionReq) (*SetCacheVersionResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetCacheVersion not implemented")
}
func (*UnimplementedHoldInfoCacheProxyServer) GetCacheVersion(ctx context.Context, req *GetCacheVersionReq) (*GetCacheVersionResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCacheVersion not implemented")
}
func (*UnimplementedHoldInfoCacheProxyServer) SaveApkConfigListCache(ctx context.Context, req *SaveApkConfigListCacheReq) (*SaveApkConfigListCacheResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveApkConfigListCache not implemented")
}
func (*UnimplementedHoldInfoCacheProxyServer) GetApkConfigListCache(ctx context.Context, req *GetApkConfigListCacheReq) (*GetApkConfigListCacheResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetApkConfigListCache not implemented")
}
func (*UnimplementedHoldInfoCacheProxyServer) GetAppKeyByHoldID(ctx context.Context, req *GetAppKeyByHoldIDReq) (*GetAppKeyByHoldIDResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppKeyByHoldID not implemented")
}
func (*UnimplementedHoldInfoCacheProxyServer) SaveApkSettingDictCache(ctx context.Context, req *SaveApkSettingDictCacheReq) (*SaveApkSettingDictCacheResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveApkSettingDictCache not implemented")
}
func (*UnimplementedHoldInfoCacheProxyServer) GetApkSettingDictCache(ctx context.Context, req *GetApkSettingDictCacheReq) (*GetApkSettingDictCacheResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetApkSettingDictCache not implemented")
}
func (*UnimplementedHoldInfoCacheProxyServer) GetHoldTreePathCache(ctx context.Context, req *GetHoldTreePathCacheReq) (*GetHoldTreePathCacheResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHoldTreePathCache not implemented")
}
func (*UnimplementedHoldInfoCacheProxyServer) SetHoldTreePathCache(ctx context.Context, req *SetHoldTreePathCacheReq) (*SetHoldTreePathCacheResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetHoldTreePathCache not implemented")
}

func RegisterHoldInfoCacheProxyServer(s *grpc.Server, srv HoldInfoCacheProxyServer) {
	s.RegisterService(&_HoldInfoCacheProxy_serviceDesc, srv)
}

func _HoldInfoCacheProxy_GetApkConfigCache_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetApkConfigCacheReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HoldInfoCacheProxyServer).GetApkConfigCache(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mapgoo.paas.cap.hold.HoldInfoCacheProxy/GetApkConfigCache",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HoldInfoCacheProxyServer).GetApkConfigCache(ctx, req.(*GetApkConfigCacheReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _HoldInfoCacheProxy_UpdateApkConfigCache_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateApkConfigCacheReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HoldInfoCacheProxyServer).UpdateApkConfigCache(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mapgoo.paas.cap.hold.HoldInfoCacheProxy/UpdateApkConfigCache",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HoldInfoCacheProxyServer).UpdateApkConfigCache(ctx, req.(*UpdateApkConfigCacheReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _HoldInfoCacheProxy_DeleteCacheInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCacheInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HoldInfoCacheProxyServer).DeleteCacheInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mapgoo.paas.cap.hold.HoldInfoCacheProxy/DeleteCacheInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HoldInfoCacheProxyServer).DeleteCacheInfo(ctx, req.(*DeleteCacheInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _HoldInfoCacheProxy_GetApkConfigCacheVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetApkConfigCacheVersionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HoldInfoCacheProxyServer).GetApkConfigCacheVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mapgoo.paas.cap.hold.HoldInfoCacheProxy/GetApkConfigCacheVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HoldInfoCacheProxyServer).GetApkConfigCacheVersion(ctx, req.(*GetApkConfigCacheVersionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _HoldInfoCacheProxy_SetApkConfigCacheVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetApkConfigCacheVersionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HoldInfoCacheProxyServer).SetApkConfigCacheVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mapgoo.paas.cap.hold.HoldInfoCacheProxy/SetApkConfigCacheVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HoldInfoCacheProxyServer).SetApkConfigCacheVersion(ctx, req.(*SetApkConfigCacheVersionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _HoldInfoCacheProxy_SetCacheVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetCacheVersionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HoldInfoCacheProxyServer).SetCacheVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mapgoo.paas.cap.hold.HoldInfoCacheProxy/SetCacheVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HoldInfoCacheProxyServer).SetCacheVersion(ctx, req.(*SetCacheVersionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _HoldInfoCacheProxy_GetCacheVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCacheVersionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HoldInfoCacheProxyServer).GetCacheVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mapgoo.paas.cap.hold.HoldInfoCacheProxy/GetCacheVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HoldInfoCacheProxyServer).GetCacheVersion(ctx, req.(*GetCacheVersionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _HoldInfoCacheProxy_SaveApkConfigListCache_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveApkConfigListCacheReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HoldInfoCacheProxyServer).SaveApkConfigListCache(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mapgoo.paas.cap.hold.HoldInfoCacheProxy/SaveApkConfigListCache",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HoldInfoCacheProxyServer).SaveApkConfigListCache(ctx, req.(*SaveApkConfigListCacheReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _HoldInfoCacheProxy_GetApkConfigListCache_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetApkConfigListCacheReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HoldInfoCacheProxyServer).GetApkConfigListCache(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mapgoo.paas.cap.hold.HoldInfoCacheProxy/GetApkConfigListCache",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HoldInfoCacheProxyServer).GetApkConfigListCache(ctx, req.(*GetApkConfigListCacheReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _HoldInfoCacheProxy_GetAppKeyByHoldID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppKeyByHoldIDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HoldInfoCacheProxyServer).GetAppKeyByHoldID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mapgoo.paas.cap.hold.HoldInfoCacheProxy/GetAppKeyByHoldID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HoldInfoCacheProxyServer).GetAppKeyByHoldID(ctx, req.(*GetAppKeyByHoldIDReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _HoldInfoCacheProxy_SaveApkSettingDictCache_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveApkSettingDictCacheReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HoldInfoCacheProxyServer).SaveApkSettingDictCache(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mapgoo.paas.cap.hold.HoldInfoCacheProxy/SaveApkSettingDictCache",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HoldInfoCacheProxyServer).SaveApkSettingDictCache(ctx, req.(*SaveApkSettingDictCacheReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _HoldInfoCacheProxy_GetApkSettingDictCache_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetApkSettingDictCacheReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HoldInfoCacheProxyServer).GetApkSettingDictCache(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mapgoo.paas.cap.hold.HoldInfoCacheProxy/GetApkSettingDictCache",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HoldInfoCacheProxyServer).GetApkSettingDictCache(ctx, req.(*GetApkSettingDictCacheReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _HoldInfoCacheProxy_GetHoldTreePathCache_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHoldTreePathCacheReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HoldInfoCacheProxyServer).GetHoldTreePathCache(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mapgoo.paas.cap.hold.HoldInfoCacheProxy/GetHoldTreePathCache",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HoldInfoCacheProxyServer).GetHoldTreePathCache(ctx, req.(*GetHoldTreePathCacheReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _HoldInfoCacheProxy_SetHoldTreePathCache_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetHoldTreePathCacheReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HoldInfoCacheProxyServer).SetHoldTreePathCache(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mapgoo.paas.cap.hold.HoldInfoCacheProxy/SetHoldTreePathCache",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HoldInfoCacheProxyServer).SetHoldTreePathCache(ctx, req.(*SetHoldTreePathCacheReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _HoldInfoCacheProxy_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mapgoo.paas.cap.hold.HoldInfoCacheProxy",
	HandlerType: (*HoldInfoCacheProxyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetApkConfigCache",
			Handler:    _HoldInfoCacheProxy_GetApkConfigCache_Handler,
		},
		{
			MethodName: "UpdateApkConfigCache",
			Handler:    _HoldInfoCacheProxy_UpdateApkConfigCache_Handler,
		},
		{
			MethodName: "DeleteCacheInfo",
			Handler:    _HoldInfoCacheProxy_DeleteCacheInfo_Handler,
		},
		{
			MethodName: "GetApkConfigCacheVersion",
			Handler:    _HoldInfoCacheProxy_GetApkConfigCacheVersion_Handler,
		},
		{
			MethodName: "SetApkConfigCacheVersion",
			Handler:    _HoldInfoCacheProxy_SetApkConfigCacheVersion_Handler,
		},
		{
			MethodName: "SetCacheVersion",
			Handler:    _HoldInfoCacheProxy_SetCacheVersion_Handler,
		},
		{
			MethodName: "GetCacheVersion",
			Handler:    _HoldInfoCacheProxy_GetCacheVersion_Handler,
		},
		{
			MethodName: "SaveApkConfigListCache",
			Handler:    _HoldInfoCacheProxy_SaveApkConfigListCache_Handler,
		},
		{
			MethodName: "GetApkConfigListCache",
			Handler:    _HoldInfoCacheProxy_GetApkConfigListCache_Handler,
		},
		{
			MethodName: "GetAppKeyByHoldID",
			Handler:    _HoldInfoCacheProxy_GetAppKeyByHoldID_Handler,
		},
		{
			MethodName: "SaveApkSettingDictCache",
			Handler:    _HoldInfoCacheProxy_SaveApkSettingDictCache_Handler,
		},
		{
			MethodName: "GetApkSettingDictCache",
			Handler:    _HoldInfoCacheProxy_GetApkSettingDictCache_Handler,
		},
		{
			MethodName: "GetHoldTreePathCache",
			Handler:    _HoldInfoCacheProxy_GetHoldTreePathCache_Handler,
		},
		{
			MethodName: "SetHoldTreePathCache",
			Handler:    _HoldInfoCacheProxy_SetHoldTreePathCache_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mapgoo.paas.cap.hold.proto",
}