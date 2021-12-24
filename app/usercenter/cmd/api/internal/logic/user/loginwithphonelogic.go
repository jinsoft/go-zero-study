package user

import (
	"context"

	"it-ku/app/usercenter/cmd/api/internal/svc"
	"it-ku/app/usercenter/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type LoginWithPhoneLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginWithPhoneLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoginWithPhoneLogic {
	return LoginWithPhoneLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginWithPhoneLogic) LoginWithPhone(req types.LoginWithPhoneReq) (resp *types.LoginSuccessResp, err error) {
	// todo: add your logic here and delete this line

	return
}
