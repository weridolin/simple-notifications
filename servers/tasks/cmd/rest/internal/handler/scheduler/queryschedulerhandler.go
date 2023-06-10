package scheduler

import (
	"net/http"

	"github.com/weridolin/simple-vedio-notifications/servers/tasks/cmd/rest/internal/logic/scheduler"
	"github.com/weridolin/simple-vedio-notifications/servers/tasks/cmd/rest/internal/svc"
	"github.com/weridolin/simple-vedio-notifications/servers/tasks/cmd/rest/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func QuerySchedulerHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.QuerySchedulerReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := scheduler.NewQuerySchedulerLogic(r.Context(), svcCtx)
		resp, err := l.QueryScheduler(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
