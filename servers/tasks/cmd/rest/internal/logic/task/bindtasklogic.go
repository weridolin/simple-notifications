package task

import (
	"context"

	"github.com/weridolin/simple-vedio-notifications/servers/tasks/cmd/rest/internal/svc"
	"github.com/weridolin/simple-vedio-notifications/servers/tasks/cmd/rest/internal/types"
	"github.com/weridolin/simple-vedio-notifications/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type BindTaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBindTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BindTaskLogic {
	return &BindTaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BindTaskLogic) BindTask(req *types.BindSchedulerReq) (resp *types.BindSchedulerResp, err error) {
	err = l.svcCtx.SchedulerModel.BindTask(req.SchedulerID, req.TaskID, l.svcCtx.DB)
	if err != nil {
		return &types.BindSchedulerResp{
			BaseResponse: types.BaseResponse{
				Code: tools.InternalError.Code,
				Msg:  err.Error(),
			},
		}, nil
	}
	return &types.BindSchedulerResp{
		BaseResponse: types.BaseResponse{
			Code: 0,
			Msg:  "绑定成功",
		}, // todo: return scheduler
	}, nil
}
