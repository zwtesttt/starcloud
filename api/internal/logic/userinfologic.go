package logic

import (
	"context"
	"fmt"
	"group/pb/starcloud/rpc/group"

	"github.com/zwtesttt/starcloud/api/internal/svc"
	"github.com/zwtesttt/starcloud/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.GetUserRequest) (resp *types.GetUserResponse, err error) {
	fmt.Println("请求进来", l.svcCtx.Group)
	_, err = l.svcCtx.Group.GetUserGroup(l.ctx, &group.GetUserGroupRequest{UserId: req.UserId})
	if err != nil {
		return nil, err
	}
	fmt.Println("rpc调用完成")

	// 初始化 resp
	resp = &types.GetUserResponse{
		UserID: req.UserId,      // 从请求中获取 UserID
		Email:  "test@test.com", // 测试数据
		Name:   "test",          // 测试数据
	}
	return resp, nil
}
