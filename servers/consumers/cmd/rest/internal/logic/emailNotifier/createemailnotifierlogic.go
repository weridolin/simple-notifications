package emailNotifier

import (
	"context"

	"github.com/weridolin/simple-vedio-notifications/servers/consumers/cmd/rest/internal/svc"
	"github.com/weridolin/simple-vedio-notifications/servers/consumers/cmd/rest/internal/types"
	"github.com/weridolin/simple-vedio-notifications/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateEmailNotifierLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateEmailNotifierLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateEmailNotifierLogic {
	return &CreateEmailNotifierLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateEmailNotifierLogic) CreateEmailNotifier(req *types.CreateEmailNotifierReq) (resp *types.CreateEmailNotifierResp, err error) {
	userID := tools.GetUidFromCtx(l.ctx)
	emailNotifier, err := l.svcCtx.EmailNotifierModel.Create(int(userID), req.PWD, req.Sender, req.Content, req.Receiver, l.svcCtx.DB)
	if err != nil {
		return &types.CreateEmailNotifierResp{
			BaseResponse: types.BaseResponse{
				Code: tools.ModelRecordCreatedError.Code,
				Msg:  err.Error(),
			},
		}, nil

	}
	return &types.CreateEmailNotifierResp{
		BaseResponse: types.BaseResponse{
			Code: 0,
			Msg:  "创建成功",
		},
		Data: *types.EmailNotifier{}.FromEmailNotifierModel(emailNotifier),
	}, nil

}
