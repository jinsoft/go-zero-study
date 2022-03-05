package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"it-ku/app/identity/cmd/api/internal/config"
	"it-ku/app/identity/cmd/rpc/identity"
)

type ServiceContext struct {
	Config      config.Config
	IdentityRpc identity.Identity
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		IdentityRpc: identity.NewIdentity(zrpc.MustNewClient(c.IdentityRpcConf)),
	}
}
