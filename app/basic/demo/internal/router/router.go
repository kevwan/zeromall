package router

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest"

	"mall/app/basic/demo/internal/dao"
	"mall/app/basic/demo/internal/service"
	"mall/app/basic/demo/proto/config"
)

func RegisterHandlers(engine *rest.Server, cfg *config.Config, serverCtx *dao.ServiceContext) {
	// new:
	s := service.NewService(cfg, serverCtx)

	// register:
	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/from/:name",
				Handler: s.Outer.DemoHandler(cfg, serverCtx),
			},
		},
	)
}
