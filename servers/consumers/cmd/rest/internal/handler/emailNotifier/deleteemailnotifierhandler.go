package emailNotifier

import (
	"net/http"

	"github.com/weridolin/simple-vedio-notifications/servers/consumers/cmd/rest/internal/logic/emailNotifier"
	"github.com/weridolin/simple-vedio-notifications/servers/consumers/cmd/rest/internal/svc"
	"github.com/weridolin/simple-vedio-notifications/servers/consumers/cmd/rest/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DeleteEmailNotifierHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteEmailNotifierReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		l := emailNotifier.NewDeleteEmailNotifierLogic(r.Context(), svcCtx)
		resp, err := l.DeleteEmailNotifier(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
