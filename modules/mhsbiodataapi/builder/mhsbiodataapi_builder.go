package builder

import (
	"tracerstudy-siak-service/common/config"
	"tracerstudy-siak-service/modules/mhsbiodataapi/handler"
	"tracerstudy-siak-service/modules/mhsbiodataapi/service"

	"google.golang.org/grpc"
)

func BuildMhsBiodataApiHandler(cfg config.Config, grpcConn *grpc.ClientConn) *handler.MhsBiodataApiHandler {
	mhsbiodataapiSvc := service.NewMhsBiodataApiService(cfg)
	return handler.NewMhsBiodataHandler(cfg, mhsbiodataapiSvc)
}
