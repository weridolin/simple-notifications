package emailNotifier

import (
	"context"

	"github.com/weridolin/simple-vedio-notifications/servers/consumers/cmd/rest/internal/svc"
	"github.com/weridolin/simple-vedio-notifications/servers/consumers/cmd/rest/internal/types"
	"github.com/weridolin/simple-vedio-notifications/tools"

	// "github.com/weridolin/simple-vedio-notifications/servers/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type BindEmailNotifierToTaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBindEmailNotifierToTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BindEmailNotifierToTaskLogic {
	return &BindEmailNotifierToTaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BindEmailNotifierToTaskLogic) BindEmailNotifierToTask(req *types.BindEmailNotifierToTaskReq) (resp *types.BindEmailNotifierToTaskResp, err error) {
	user_id := tools.GetUidFromCtx(l.ctx)
	err = l.svcCtx.EmailNotifierModel.BindEmailNotifierToTask(int(user_id), req.TaskId, req.EmailNotifierId, l.svcCtx.DB)
	if err != nil {
		return &types.BindEmailNotifierToTaskResp{
			BaseResponse: types.BaseResponse{
				Code: tools.InternalError.Code,
				Msg:  err.Error(),
			},
		}, nil
	}
	return &types.BindEmailNotifierToTaskResp{
		BaseResponse: types.BaseResponse{
			Code: 0,
			Msg:  "绑定成功",
		},
	}, nil
}
