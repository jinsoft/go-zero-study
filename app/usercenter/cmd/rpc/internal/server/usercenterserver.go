// Code generated by goctl. DO NOT EDIT!
// Source: usercenter.proto

package server

import (
	"context"

	"it-ku/app/usercenter/cmd/rpc/internal/logic"
	"it-ku/app/usercenter/cmd/rpc/internal/svc"
	"it-ku/app/usercenter/cmd/rpc/pb"
)

type UsercenterServer struct {
	svcCtx *svc.ServiceContext
}

func NewUsercenterServer(svcCtx *svc.ServiceContext) *UsercenterServer {
	return &UsercenterServer{
		svcCtx: svcCtx,
	}
}

func (s *UsercenterServer) RegisterWithEmail(ctx context.Context, in *pb.RegisterWithEmailReq) (*pb.LoginSuccessResp, error) {
	l := logic.NewRegisterWithEmailLogic(ctx, s.svcCtx)
	return l.RegisterWithEmail(in)
}

func (s *UsercenterServer) LoginWithEmail(ctx context.Context, in *pb.LoginWithEmailReq) (*pb.LoginSuccessResp, error) {
	l := logic.NewLoginWithEmailLogic(ctx, s.svcCtx)
	return l.LoginWithEmail(in)
}

func (s *UsercenterServer) GetUserInfo(ctx context.Context, in *pb.GetUserInfoReq) (*pb.GetUserInfoResp, error) {
	l := logic.NewGetUserInfoLogic(ctx, s.svcCtx)
	return l.GetUserInfo(in)
}