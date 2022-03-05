package user

import (
	"context"

	"it-ku/app/usercenter/cmd/api/internal/svc"
	"it-ku/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdateLogic {
	return UpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateLogic) Update(req types.UpdateReq) (resp *types.UserInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
