package user

import (
	"net/http"

	"github.com/weridolin/simple-vedio-notifications/servers/users/cmd/rest/internal/logic/user"
	"github.com/weridolin/simple-vedio-notifications/servers/users/cmd/rest/internal/svc"
	"github.com/weridolin/simple-vedio-notifications/servers/users/cmd/rest/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func TokenValidateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ValidateTokenReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewTokenValidateLogic(r.Context(), svcCtx)
		resp, err := l.TokenValidate(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			w.Header().Set("X-Forwarded-User", resp.UserId)
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
