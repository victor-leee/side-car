// Code generated by protoc-gen-go-scrpc. DO NOT EDIT.
// Path proto/config-backend.proto
// This is the client side implementation
// If you would like to generate the server side, please add --go-scrpc_opt=server=true to your protoc command

package config_backend

import (
	context "context"
	scrpc "github.com/victor-leee/scrpc"
)

var client scrpc.Client

func init() {
	client = scrpc.NewClient()
}

type ConfigBackendService interface {
	GetConfig(ctx context.Context, req *GetConfigRequest) (*GetConfigResponse, error)
	PutConfig(ctx context.Context, req *PutConfigRequest) (*PutConfigResponse, error)
	GetAllKeys(ctx context.Context, req *GetAllKeysRequest) (*GetAllKeysResponse, error)
}

type ConfigBackendServiceImpl struct {
}

func (*ConfigBackendServiceImpl) GetConfig(ctx context.Context, req *GetConfigRequest) (*GetConfigResponse, error) {
	resp := &GetConfigResponse{}
	err := client.UnaryRPCRequest(&scrpc.RequestContext{
		Ctx:           ctx,
		Req:           req,
		ReqService:    "config-backend-victor-leee-github-com",
		ReqMethod:     "GetConfig",
		SenderService: "side-car",
		Resp:          resp,
	})
	return resp, err
}
func (*ConfigBackendServiceImpl) PutConfig(ctx context.Context, req *PutConfigRequest) (*PutConfigResponse, error) {
	resp := &PutConfigResponse{}
	err := client.UnaryRPCRequest(&scrpc.RequestContext{
		Ctx:           ctx,
		Req:           req,
		ReqService:    "config-backend-victor-leee-github-com",
		ReqMethod:     "PutConfig",
		SenderService: "side-car",
		Resp:          resp,
	})
	return resp, err
}
func (*ConfigBackendServiceImpl) GetAllKeys(ctx context.Context, req *GetAllKeysRequest) (*GetAllKeysResponse, error) {
	resp := &GetAllKeysResponse{}
	err := client.UnaryRPCRequest(&scrpc.RequestContext{
		Ctx:           ctx,
		Req:           req,
		ReqService:    "config-backend-victor-leee-github-com",
		ReqMethod:     "GetAllKeys",
		SenderService: "side-car",
		Resp:          resp,
	})
	return resp, err
}
