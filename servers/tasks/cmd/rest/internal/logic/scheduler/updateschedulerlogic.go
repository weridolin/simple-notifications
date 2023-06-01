package scheduler

import (
	"context"

	"github.com/weridolin/simple-vedio-notifications/servers/tasks/cmd/rest/internal/svc"
	"github.com/weridolin/simple-vedio-notifications/servers/tasks/cmd/rest/internal/types"

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
	// todo: add your logic here and delete this line

	return
}
