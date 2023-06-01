package task

import (
	"context"

	"github.com/weridolin/simple-vedio-notifications/servers/tasks/cmd/rest/internal/svc"
	"github.com/weridolin/simple-vedio-notifications/servers/tasks/cmd/rest/internal/types"

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
	// todo: add your logic here and delete this line

	return
}
