package task

import (
	"context"

	"github.com/weridolin/simple-vedio-notifications/servers/tasks/cmd/rest/internal/svc"
	"github.com/weridolin/simple-vedio-notifications/servers/tasks/cmd/rest/internal/types"
	"github.com/weridolin/simple-vedio-notifications/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateTaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTaskLogic {
	return &UpdateTaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateTaskLogic) UpdateTask(req *types.UpdateTaskReq) (resp *types.UpdateTaskResp, err error) {
	id := req.ID
	userId := tools.GetUidFromCtx(l.ctx)
	task, err := l.svcCtx.TaskModel.Query(map[string]interface{}{"id": id, "user_id": userId}, l.svcCtx.DB)
	if err != nil {
		return &types.UpdateTaskResp{
			CreateTaskResp: types.CreateTaskResp{
				BaseResponse: types.BaseResponse{
					Code: tools.ModelRecordUpdatedError.Code,
					Msg:  err.Error(),
				},
			},
		}, nil
	}
	task.Ups = req.Ups
	task.Platform = req.Platform
	task.Name = req.Name
	task.Description = req.Description
	err = l.svcCtx.TaskModel.Update(task, l.svcCtx.DB)
	if err != nil {
		return &types.UpdateTaskResp{
			CreateTaskResp: types.CreateTaskResp{
				BaseResponse: types.BaseResponse{
					Code: tools.ModelRecordUpdatedError.Code,
					Msg:  err.Error(),
				},
			},
		}, nil
	}
	return &types.UpdateTaskResp{
		CreateTaskResp: types.CreateTaskResp{
			BaseResponse: types.BaseResponse{
				Code: 0,
				Msg:  "更新成功",
			},
			Data: *types.Task{}.FromTaskModel(task),
		},
	}, nil
}
