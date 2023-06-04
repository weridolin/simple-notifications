package emailNotifier

import (
	"context"

	"github.com/weridolin/simple-vedio-notifications/servers/consumers/cmd/rest/internal/svc"
	"github.com/weridolin/simple-vedio-notifications/servers/consumers/cmd/rest/internal/types"
	"github.com/weridolin/simple-vedio-notifications/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteEmailNotifierLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteEmailNotifierLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteEmailNotifierLogic {
	return &DeleteEmailNotifierLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteEmailNotifierLogic) DeleteEmailNotifier(req *types.DeleteEmailNotifierReq) (resp *types.DeleteEmailNotifierResp, err error) {
	userId := tools.GetUidFromCtx(l.ctx)
	_, err = l.svcCtx.EmailNotifierModel.Query(map[string]interface{}{"user_id": userId, "id": req.Id}, l.svcCtx.DB)
	if err != nil {
		return &types.DeleteEmailNotifierResp{
			BaseResponse: types.BaseResponse{
				Code: tools.ModelRecordDeletedError.Code,
				Msg:  err.Error(),
			},
		}, nil
	}
	l.svcCtx.EmailNotifierModel.Delete(req.Id, l.svcCtx.DB)
	return &types.DeleteEmailNotifierResp{
		BaseResponse: types.BaseResponse{
			Code: 0,
			Msg:  "删除成功",
		},
	}, nil
}
