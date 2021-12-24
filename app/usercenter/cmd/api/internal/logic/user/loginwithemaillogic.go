package user

import (
	"context"

	"it-ku/app/usercenter/cmd/api/internal/svc"
	"it-ku/app/usercenter/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type LoginWithEmailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginWithEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoginWithEmailLogic {
	return LoginWithEmailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginWithEmailLogic) LoginWithEmail(req types.LoginWithEmailReq) (resp *types.LoginSuccessResp, err error) {
	// todo: add your logic here and delete this line

	return
}
