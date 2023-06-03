package task

import (
	"context"

	"github.com/weridolin/simple-vedio-notifications/servers/tasks/cmd/rest/internal/svc"
	"github.com/weridolin/simple-vedio-notifications/servers/tasks/cmd/rest/internal/types"
	"github.com/weridolin/simple-vedio-notifications/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateTaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateTaskLogic {
	return &CreateTaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateTaskLogic) CreateTask(req *types.CreateTaskReq) (resp *types.CreateTaskResp, err error) {
	userId := tools.GetUidFromCtx(l.ctx)
	task, err := l.svcCtx.TaskModel.Create(int(userId), req.Ups, req.Platform, req.Name, req.Description, l.svcCtx.DB)
	if err != nil {
		return &types.CreateTaskResp{
			BaseResponse: types.BaseResponse{
				Code: tools.ModelRecordCreatedError.Code,
				Msg:  err.Error(),
			},
		}, nil
	}
	return &types.CreateTaskResp{
		BaseResponse: types.BaseResponse{
			Code: 0,
			Msg:  "创建成功",
		},
		Data: *types.Task{}.FromTaskModel(task),
	}, nil
}
