package scheduler

import (
	"context"

	"github.com/weridolin/simple-vedio-notifications/servers/tasks/cmd/rest/internal/svc"
	"github.com/weridolin/simple-vedio-notifications/servers/tasks/cmd/rest/internal/types"
	"github.com/weridolin/simple-vedio-notifications/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateSchedulerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateSchedulerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateSchedulerLogic {
	return &CreateSchedulerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateSchedulerLogic) CreateScheduler(req *types.CreateSchedulerReq) (resp *types.CreateSchedulerResp, err error) {
	userID := tools.GetUidFromCtx(l.ctx)
	scheduler, err := l.svcCtx.SchedulerModel.Create(int(userID), req.Period, req.Platform, req.Name, req.Description, l.svcCtx.DB)
	if err != nil {
		return &types.CreateSchedulerResp{
			BaseResponse: types.BaseResponse{
				Code: tools.ModelRecordCreatedError.Code,
				Msg:  err.Error(),
			},
		}, nil
	}

	return &types.CreateSchedulerResp{
		BaseResponse: types.BaseResponse{
			Code: 0,
			Msg:  "创建成功",
		},
		Data: types.Scheduler{}.FromSchedulerModel(scheduler),
	}, nil
}
