package emailNotifier

import (
	"context"
	"fmt"

	"github.com/weridolin/simple-vedio-notifications/servers/consumers/cmd/rest/internal/svc"
	"github.com/weridolin/simple-vedio-notifications/servers/consumers/cmd/rest/internal/types"
	"github.com/weridolin/simple-vedio-notifications/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryEmailNotifierLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryEmailNotifierLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryEmailNotifierLogic {
	return &QueryEmailNotifierLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryEmailNotifierLogic) QueryEmailNotifier(req *types.QueryEmailNotifierReq) (resp *types.QueryEmailNotifierResp, err error) {
	userId := tools.GetUidFromCtx(l.ctx)
	emailNotifierList, err := l.svcCtx.EmailNotifierModel.QueryAll(map[string]interface{}{"user_id": int(userId)}, req.Page, req.Size, l.svcCtx.DB)
	if err != nil {
		return &types.QueryEmailNotifierResp{
			BaseResponse: types.BaseResponse{
				Code: tools.ModelQueryError.Code,
				Msg:  err.Error(),
			},
		}, nil
	}
	fmt.Println(emailNotifierList, userId)
	return &types.QueryEmailNotifierResp{
		BaseResponse: types.BaseResponse{
			Code: 0,
			Msg:  "查询成功",
		},
		Data: types.EmailNotifier{}.FromEmailNotifierModels(emailNotifierList),
	}, nil

}
