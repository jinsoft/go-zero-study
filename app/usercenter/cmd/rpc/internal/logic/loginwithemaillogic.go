package logic

import (
	"context"

	"it-ku/app/usercenter/cmd/rpc/internal/svc"
	"it-ku/app/usercenter/cmd/rpc/pb"

	"github.com/tal-tech/go-zero/core/logx"
)

type LoginWithEmailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginWithEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginWithEmailLogic {
	return &LoginWithEmailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginWithEmailLogic) LoginWithEmail(in *pb.LoginWithEmailReq) (*pb.LoginSuccessResp, error) {
	// todo: add your logic here and delete this line

	return &pb.LoginSuccessResp{}, nil
}
