package scheduler

import (
	"context"

	"github.com/weridolin/simple-vedio-notifications/servers/tasks/cmd/rest/internal/svc"
	"github.com/weridolin/simple-vedio-notifications/servers/tasks/cmd/rest/internal/types"
	"github.com/weridolin/simple-vedio-notifications/tools"

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
	userID := tools.GetUidFromCtx(l.ctx)
	_, err = l.svcCtx.SchedulerModel.Query(map[string]interface{}{"id": req.ID, "user_id": userID}, l.svcCtx.DB)
	if err != nil {
		return &types.DeleteSchedulerResp{
			BaseResponse: types.BaseResponse{
				Code: tools.ModelRecordDeletedError.Code,
				Msg:  err.Error(),
			},
		}, nil
	}
	l.svcCtx.SchedulerModel.Delete(req.ID, l.svcCtx.DB)
	return &types.DeleteSchedulerResp{
		BaseResponse: types.BaseResponse{
			Code: 0,
			Msg:  "删除成功",
		},
	}, nil
}
