package scheduler

import (
	"context"

	"github.com/weridolin/simple-vedio-notifications/servers/tasks/cmd/rest/internal/svc"
	"github.com/weridolin/simple-vedio-notifications/servers/tasks/cmd/rest/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteSchedulerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteSchedulerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSchedulerLogic {
	return &DeleteSchedulerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteSchedulerLogic) DeleteScheduler(req *types.DeleteSchedulerReq) (resp *types.DeleteSchedulerResp, err error) {
	// todo: add your logic here and delete this line

	return
}
