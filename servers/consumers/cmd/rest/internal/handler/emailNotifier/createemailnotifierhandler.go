package emailNotifier

import (
	"net/http"

	"github.com/weridolin/simple-vedio-notifications/servers/consumers/cmd/rest/internal/logic/emailNotifier"
	"github.com/weridolin/simple-vedio-notifications/servers/consumers/cmd/rest/internal/svc"
	"github.com/weridolin/simple-vedio-notifications/servers/consumers/cmd/rest/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CreateEmailNotifierHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateEmailNotifierReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := emailNotifier.NewCreateEmailNotifierLogic(r.Context(), svcCtx)
		resp, err := l.CreateEmailNotifier(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
