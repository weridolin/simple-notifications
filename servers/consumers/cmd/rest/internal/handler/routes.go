// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	emailNotifier "github.com/weridolin/simple-vedio-notifications/servers/consumers/cmd/rest/internal/handler/emailNotifier"
	"github.com/weridolin/simple-vedio-notifications/servers/consumers/cmd/rest/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/emailNotifier",
				Handler: emailNotifier.CreateEmailNotifierHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/emailNotifier/:id",
				Handler: emailNotifier.DeleteEmailNotifierHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/emailNotifier/:id",
				Handler: emailNotifier.UpdateEmailNotifierHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/emailNotifier",
				Handler: emailNotifier.QueryEmailNotifierHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/emailNotifier/bind",
				Handler: emailNotifier.BindEmailNotifierToTaskHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/api/v1"),
	)
}
