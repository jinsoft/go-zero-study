package user

import (
	"context"

	"it-ku/app/usercenter/cmd/api/internal/svc"
	"it-ku/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterWithEmailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterWithEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) RegisterWithEmailLogic {
	return RegisterWithEmailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterWithEmailLogic) RegisterWithEmail(req types.EmailRegisterReq) (resp *types.LoginSuccessResp, err error) {
	// todo: add your logic here and delete this line

	return
}
