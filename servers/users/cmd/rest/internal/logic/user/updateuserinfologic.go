package user

import (
	"context"

	"github.com/weridolin/simple-vedio-notifications/servers/users/cmd/rest/internal/svc"
	"github.com/weridolin/simple-vedio-notifications/servers/users/cmd/rest/internal/types"
	"github.com/weridolin/simple-vedio-notifications/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserInfoLogic {
	return &UpdateUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserInfoLogic) UpdateUserInfo(req *types.UpdateUserInfoRep) (resp *types.UpdateUserInfoResp, err error) {
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
	user.Age = req.Age
	user.Avatar = req.Avatar
	user.Username = req.Username
	user.Email = req.Email
	user.Phone = req.Phone
	user.Gender = req.Gender
	err = l.svcCtx.UserModel.Update(user, l.svcCtx.DB)
	if err != nil {
		return &types.UpdateUserInfoResp{
			BaseResponse: types.BaseResponse{
				Code: tools.ModelRecordDeletedError.Code,
				Msg:  err.Error(),
			},
		}, nil
	}
	return &types.UpdateUserInfoResp{
		BaseResponse: types.BaseResponse{
			Code: 0,
			Msg:  "更新成功",
		},
		Data: *types.UserInfo{}.FromUserModel(user),
	}, nil
}
