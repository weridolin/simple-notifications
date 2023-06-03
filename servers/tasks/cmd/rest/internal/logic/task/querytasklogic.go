package task

import (
	"context"
	"fmt"

	"github.com/weridolin/simple-vedio-notifications/servers/tasks/cmd/rest/internal/svc"
	"github.com/weridolin/simple-vedio-notifications/servers/tasks/cmd/rest/internal/types"
	"github.com/weridolin/simple-vedio-notifications/tools"

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
	userId := tools.GetUidFromCtx(l.ctx)
	fmt.Println(userId)
	tasks, err := l.svcCtx.TaskModel.FullQueryTasks(map[string]interface{}{"user_id": userId}, req.Page, req.Size, l.svcCtx.DB)
	if err != nil {
		return &types.QueryTaskResp{
			BaseResponse: types.BaseResponse{
				Code: tools.ModelQueryError.Code,
				Msg:  err.Error(),
			},
		}, nil
	}
	return &types.QueryTaskResp{
		BaseResponse: types.BaseResponse{
			Code: 0,
			Msg:  "查询成功",
		},
		Data: types.Task{}.FromTaskModels(tasks),
	}, nil
}
