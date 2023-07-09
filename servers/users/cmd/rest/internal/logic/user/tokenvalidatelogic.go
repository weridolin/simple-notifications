package user

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/weridolin/simple-vedio-notifications/servers/users/cmd/rest/internal/svc"
	"github.com/weridolin/simple-vedio-notifications/servers/users/cmd/rest/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TokenValidateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTokenValidateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TokenValidateLogic {
	return &TokenValidateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TokenValidateLogic) TokenValidate(req *types.ValidateTokenReq) (resp *types.ValidateResp, err error) {
	fmt.Println("token validate logic", req.Authorization)
	userID := l.ctx.Value("id")
	err = TokenUnregister(req.Authorization)
	if err != nil {
		return nil, err
	} else {
		return &types.ValidateResp{
			UserId: userID.(json.Number).String(),
		}, nil
	}
}

// todo token是否被注销
func TokenUnregister(token string) error {
	return nil
}
