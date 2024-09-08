package svc

import (
	"api/internal/config"
	"github.com/zeromicro/go-zero/zrpc"
	"group/groupservice"
)

type ServiceContext struct {
	Config config.Config
	Group  groupservice.GroupService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Group:  groupservice.NewGroupService(zrpc.MustNewClient(c.GroupRpc)),
	}
}
