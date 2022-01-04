package logic

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"it-ku/common/globalkey"
	"it-ku/common/xerr"

	"it-ku/app/identity/cmd/rpc/internal/svc"
	"it-ku/app/identity/cmd/rpc/pb"

	"github.com/tal-tech/go-zero/core/logx"
)

var ErrClearTokenError = xerr.NewErrMsg("退出token失败")

type ClearTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewClearTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClearTokenLogic {
	return &ClearTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  清理token, 只让用户服务访问
func (l *ClearTokenLogic) ClearToken(in *pb.ClearTokenReq) (*pb.ClearTokenResp, error) {
	userTokenKey := fmt.Sprintf(globalkey.UserTokenKey, in.UserId)
	if _, err := l.svcCtx.RedisClient.Del(userTokenKey); err != nil {
		return nil, errors.Wrapf(ErrClearTokenError, "userId:%d,err:%v", in.UserId, err)
	}
	return &pb.ClearTokenResp{
		Ok: true,
	}, nil
}
