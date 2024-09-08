package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zwtesttt/starcloud/api/internal/config"
	"github.com/zwtesttt/starcloud/rpc/group/groupservice"
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
