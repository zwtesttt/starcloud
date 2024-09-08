package logic

import (
	"context"
	"fmt"

	"group/internal/svc"
	"group/pb/starcloud/rpc/group"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserGroupLogic {
	return &GetUserGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserGroupLogic) GetUserGroup(in *group.GetUserGroupRequest) (*group.GetUserGroupResponse, error) {
	fmt.Println("GetUserGroup请求进来")
	if in == nil {
		return nil, fmt.Errorf("参数错误")
	}
	fmt.Println("in.UserId =>", in.UserId)

	return &group.GetUserGroupResponse{
		GroupIds: []string{"1", "2"},
	}, nil
}
