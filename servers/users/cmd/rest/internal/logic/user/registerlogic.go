package user

import (
	"context"

	"github.com/weridolin/simple-vedio-notifications/servers/users/cmd/rest/internal/svc"
	"github.com/weridolin/simple-vedio-notifications/servers/users/cmd/rest/internal/types"
	"github.com/weridolin/simple-vedio-notifications/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	if user, err := l.svcCtx.UserModel.Create(req.Username, req.Email, req.Password, l.svcCtx.DB); err != nil {
		return &types.RegisterResp{
			BaseResponse: types.BaseResponse{
				Code: tools.ModelAlreadyExist.Code,
				Msg:  err.Error(),
			},
		}, nil
	} else {
		return &types.RegisterResp{
			BaseResponse: types.BaseResponse{
				Code: 0,
				Msg:  "注册成功",
			},
			Data: *types.UserInfo{}.FromUserModel(*user),
		}, nil
	}
}
