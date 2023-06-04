package emailNotifier

import (
	"context"

	"github.com/weridolin/simple-vedio-notifications/servers/consumers/cmd/rest/internal/svc"
	"github.com/weridolin/simple-vedio-notifications/servers/consumers/cmd/rest/internal/types"

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
	// todo: add your logic here and delete this line

	return
}
