package emailNotifier

import (
	"context"

	"github.com/weridolin/simple-vedio-notifications/servers/consumers/cmd/rest/internal/svc"
	"github.com/weridolin/simple-vedio-notifications/servers/consumers/cmd/rest/internal/types"
	"github.com/weridolin/simple-vedio-notifications/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateEmailNotifierLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateEmailNotifierLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateEmailNotifierLogic {
	return &UpdateEmailNotifierLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateEmailNotifierLogic) UpdateEmailNotifier(req *types.UpdateEmailNotifierReq) (resp *types.UpdateEmailNotifierResp, err error) {
	userId := tools.GetUidFromCtx(l.ctx)
	emailNotifier, err := l.svcCtx.EmailNotifierModel.Query(map[string]interface{}{"user_id": userId, "id": req.ID}, l.svcCtx.DB)
	if err != nil {
		return &types.UpdateEmailNotifierResp{
			BaseResponse: types.BaseResponse{
				Code: tools.ModelRecordUpdatedError.Code,
				Msg:  err.Error(),
			},
		}, nil
	}
	return &types.UpdateEmailNotifierResp{
		BaseResponse: types.BaseResponse{
			Code: 0,
			Msg:  "更新成功",
		},
		Data: *types.EmailNotifier{}.FromEmailNotifierModel(emailNotifier),
	}, nil
}
