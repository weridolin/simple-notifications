package user

import (
	"context"

	"github.com/weridolin/simple-vedio-notifications/servers/users/cmd/rest/internal/svc"
	"github.com/weridolin/simple-vedio-notifications/servers/users/cmd/rest/internal/types"
	"github.com/weridolin/simple-vedio-notifications/tools"

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

func (l *UserInfoDetailLogic) UserInfoDetail() (resp *types.UpdateUserInfoResp, err error) {
	userID := l.ctx.Value("userId")
	user, err := l.svcCtx.UserModel.QueryUser(map[string]interface{}{"id": userID}, l.svcCtx.DB)
	if err != nil {
		return &types.UpdateUserInfoResp{
			BaseResponse: types.BaseResponse{
				Code: tools.ModelRecordNotFound.Code,
				Msg:  "用户不存在",
			},
		}, nil
	}
	return &types.UpdateUserInfoResp{
		BaseResponse: types.BaseResponse{
			Code: 0,
			Msg:  "查询成功",
		},
		Data: *types.UserInfo{}.FromUserModel(user),
	}, nil
}
