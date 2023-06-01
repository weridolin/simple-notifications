package task

import (
	"context"

	"github.com/weridolin/simple-vedio-notifications/servers/tasks/cmd/rest/internal/svc"
	"github.com/weridolin/simple-vedio-notifications/servers/tasks/cmd/rest/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryTaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryTaskLogic {
	return &QueryTaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryTaskLogic) QueryTask(req *types.QueryTaskReq) (resp *types.QueryTaskResp, err error) {
	// todo: add your logic here and delete this line

	return
}
