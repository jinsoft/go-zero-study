package user

import (
	"context"

	"it-ku/app/usercenter/cmd/api/internal/svc"
	"it-ku/app/usercenter/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type RegisterWithPhoneLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterWithPhoneLogic(ctx context.Context, svcCtx *svc.ServiceContext) RegisterWithPhoneLogic {
	return RegisterWithPhoneLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterWithPhoneLogic) RegisterWithPhone(req types.MobileRegisterReq) (resp *types.LoginSuccessResp, err error) {
	// todo: add your logic here and delete this line

	return
}
