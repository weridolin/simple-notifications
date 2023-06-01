package scheduler

import (
	"context"

	"github.com/weridolin/simple-vedio-notifications/servers/tasks/cmd/rest/internal/svc"
	"github.com/weridolin/simple-vedio-notifications/servers/tasks/cmd/rest/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QuerySchedulerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQuerySchedulerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QuerySchedulerLogic {
	return &QuerySchedulerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QuerySchedulerLogic) QueryScheduler(req *types.QuerySchedulerReq) (resp *types.QuerySchedulerResp, err error) {
	// todo: add your logic here and delete this line

	return
}
