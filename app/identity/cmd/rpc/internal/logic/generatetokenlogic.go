package logic

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"it-ku/common/ctxdata"
	"it-ku/common/globalkey"
	"it-ku/common/xerr"
	"time"

	"it-ku/app/identity/cmd/rpc/internal/svc"
	"it-ku/app/identity/cmd/rpc/pb"

	"github.com/dgrijalva/jwt-go"
	"github.com/zeromicro/go-zero/core/logx"
)

var ErrGenerateTokenError = xerr.NewErrMsg("生成token失败")

type GenerateTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGenerateTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateTokenLogic {
	return &GenerateTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  生成token, 只让用户服务访问
func (l *GenerateTokenLogic) GenerateToken(in *pb.GenerateTokenReq) (*pb.GenerateTokenResp, error) {
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.JwtAuth.AccessExpire
	accessToken, err := l.getJwtToken(l.svcCtx.Config.JwtAuth.AccessSecret, now, accessExpire, in.UserId)
	if err != nil {
		return nil, errors.Wrapf(ErrGenerateTokenError, "getJwtToken err userId:%d , err:%v", in.UserId, err)
	}

	userTokenKey := fmt.Sprintf(globalkey.UserTokenKey, in.UserId)
	err = l.svcCtx.RedisClient.Setex(userTokenKey, accessToken, int(accessExpire))
	if err != nil {
		return nil, errors.Wrapf(ErrGenerateTokenError, "SetnxEx err userId:%d, err:%v", in.UserId, err)
	}
	return &pb.GenerateTokenResp{
		AccessToken:  accessToken,
		AccessExpire: now + accessExpire,
		RefreshAfter: now + accessExpire/2,
	}, nil
}

func (l *GenerateTokenLogic) getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {

	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims[ctxdata.CtxKeyJwtUserId] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
