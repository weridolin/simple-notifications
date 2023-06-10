package task

import (
	"context"

	"github.com/weridolin/simple-vedio-notifications/servers/tasks/cmd/rest/internal/svc"
	"github.com/weridolin/simple-vedio-notifications/servers/tasks/cmd/rest/internal/types"
	"github.com/weridolin/simple-vedio-notifications/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteTaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteTaskLogic {
	return &DeleteTaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteTaskLogic) DeleteTask(req *types.DeleteTaskReq) (resp *types.DeleteTaskResp, err error) {
	userId := tools.GetUidFromCtx(l.ctx)
	_, err = l.svcCtx.TaskModel.Query(map[string]interface{}{"id": req.ID, "user_id": userId}, l.svcCtx.DB)
	if err != nil {
		return &types.DeleteTaskResp{
			BaseResponse: types.BaseResponse{
				Code: tools.ModelRecordDeletedError.Code,
				Msg:  err.Error(),
			},
		}, nil
	}
	err = l.svcCtx.TaskModel.Delete(req.ID, l.svcCtx.DB)
	if err != nil {
		return &types.DeleteTaskResp{
			BaseResponse: types.BaseResponse{
				Code: tools.ModelRecordDeletedError.Code,
				Msg:  err.Error(),
			},
		}, nil
	}
	return &types.DeleteTaskResp{
		BaseResponse: types.BaseResponse{
			Code: 0,
			Msg:  "删除成功",
		},
	}, nil
}
