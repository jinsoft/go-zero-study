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

type ValidateTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

var ValidateTokenError = xerr.NewErrCode(xerr.TokenExpireError)

func NewValidateTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ValidateTokenLogic {
	return &ValidateTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ValidateTokenLogic) ValidateToken(in *pb.ValidateTokenReq) (*pb.ValidateTokenResp, error) {

	userTokenKey := fmt.Sprintf(globalkey.UserTokenKey, in.UserId)
	dbToken, err := l.svcCtx.RedisClient.Get(userTokenKey)
	if err != nil {
		return nil, errors.Wrapf(ValidateTokenError, "ValidateToken RedisClient Get userId:%d ,err:%v", in.UserId, err)
	}

	if dbToken != in.Token {
		return nil, errors.Wrapf(ValidateTokenError, "ValidateToken is invalid  userId:%d , token:%s , dbToken:%s", in.UserId, in.Token, dbToken)
	}

	return &pb.ValidateTokenResp{
		Ok: true,
	}, nil
}
