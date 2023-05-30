package user

import (
	"context"

	"github.com/weridolin/simple-vedio-notifications/servers/users/cmd/rest/internal/svc"
	"github.com/weridolin/simple-vedio-notifications/servers/users/cmd/rest/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoDetailLogic {
	return &UserInfoDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoDetailLogic) UserInfoDetail() (resp *types.UserInfo, err error) {
	// todo: add your logic here and delete this line

	return
}
