package logic

import (
	"context"
	"github.com/pkg/errors"
	"it-ku/app/usercenter/model"
	"it-ku/common/xerr"

	"it-ku/app/usercenter/cmd/rpc/internal/svc"
	"it-ku/app/usercenter/cmd/rpc/pb"

	"github.com/tal-tech/go-zero/core/logx"
)

type RegisterWithEmailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

var ErrUserAlreadyRegisterError = xerr.NewErrMsg("该用户已被注册")

func NewRegisterWithEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterWithEmailLogic {
	return &RegisterWithEmailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterWithEmailLogic) RegisterWithEmail(in *pb.RegisterWithEmailReq) (*pb.LoginSuccessResp, error) {
	user, err := l.svcCtx.UserModel.FindOneByEmail(in.Email)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "email:%s,err:%v", in.Email, err)
	}
	if user != nil {
		return nil, errors.Wrapf(ErrUserAlreadyRegisterError, "用户已经存在 email:%s,err:%v", in.Email, err)
	}
	return &pb.LoginSuccessResp{}, nil
}
