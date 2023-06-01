package scheduler

import (
	"context"

	"github.com/weridolin/simple-vedio-notifications/servers/tasks/cmd/rest/internal/svc"
	"github.com/weridolin/simple-vedio-notifications/servers/tasks/cmd/rest/internal/types"

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
	// todo: add your logic here and delete this line

	return
}
