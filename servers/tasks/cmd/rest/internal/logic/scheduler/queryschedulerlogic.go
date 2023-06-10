package scheduler

import (
	"context"

	"github.com/weridolin/simple-vedio-notifications/servers/tasks/cmd/rest/internal/svc"
	"github.com/weridolin/simple-vedio-notifications/servers/tasks/cmd/rest/internal/types"
	"github.com/weridolin/simple-vedio-notifications/tools"

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
	userId := tools.GetUidFromCtx(l.ctx)
	res, err := l.svcCtx.SchedulerModel.FullQuery(map[string]interface{}{"user_id": userId}, req.Page, req.Size, l.svcCtx.DB)
	if err != nil {
		return &types.QuerySchedulerResp{
			BaseResponse: types.BaseResponse{
				Code: tools.ModelQueryError.Code,
				Msg:  err.Error(),
			},
		}, nil
	}
	return &types.QuerySchedulerResp{
		BaseResponse: types.BaseResponse{
			Code: 0,
			Msg:  "查询成功",
		},
		Data: types.Scheduler{}.FromSchedulerModels(res),
	}, nil
}
