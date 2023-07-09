package user

import (
	"context"

	"github.com/weridolin/simple-vedio-notifications/servers/users/cmd/rest/internal/svc"
	"github.com/weridolin/simple-vedio-notifications/servers/users/cmd/rest/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TokenRefreshLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTokenRefreshLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TokenRefreshLogic {
	return &TokenRefreshLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TokenRefreshLogic) TokenRefresh(req *types.RefreshTokenReq) (resp *types.RefreshTokenResp, err error) {
	// todo: add your logic here and delete this line

	return
}
