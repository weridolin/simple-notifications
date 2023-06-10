package user

import (
	"net/http"

	"github.com/weridolin/simple-vedio-notifications/servers/users/cmd/rest/internal/logic/user"
	"github.com/weridolin/simple-vedio-notifications/servers/users/cmd/rest/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserInfoDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewUserInfoDetailLogic(r.Context(), svcCtx)
		resp, err := l.UserInfoDetail()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
