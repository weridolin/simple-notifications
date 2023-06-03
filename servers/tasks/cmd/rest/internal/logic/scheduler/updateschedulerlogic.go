package scheduler

import (
	"context"

	"github.com/weridolin/simple-vedio-notifications/servers/tasks/cmd/rest/internal/svc"
	"github.com/weridolin/simple-vedio-notifications/servers/tasks/cmd/rest/internal/types"
	"github.com/weridolin/simple-vedio-notifications/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSchedulerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateSchedulerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSchedulerLogic {
	return &UpdateSchedulerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateSchedulerLogic) UpdateScheduler(req *types.UpdateSchedulerReq) (resp *types.UpdateSchedulerResp, err error) {
	userId := tools.GetUidFromCtx(l.ctx)
	scheduler, err := l.svcCtx.SchedulerModel.Query(map[string]interface{}{"id": req.ID, "user_id": userId}, l.svcCtx.DB)
	if err != nil {
		return &types.UpdateSchedulerResp{
			CreateSchedulerResp: types.CreateSchedulerResp{
				BaseResponse: types.BaseResponse{
					Code: tools.ModelRecordUpdatedError.Code,
					Msg:  err.Error(),
				},
			},
		}, nil
	}
	scheduler.Name = req.Name
	scheduler.Description = req.Description
	scheduler.Period = req.Period
	scheduler.Platform = req.Platform
	return &types.UpdateSchedulerResp{
		CreateSchedulerResp: types.CreateSchedulerResp{
			BaseResponse: types.BaseResponse{
				Code: 0,
				Msg:  "更新成功",
			},
			Data: types.Scheduler{}.FromSchedulerModel(scheduler),
		},
	}, nil
}
