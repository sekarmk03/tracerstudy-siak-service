package mhsbiodata

import (
	"tracerstudy-siak-service/common/config"
	"tracerstudy-siak-service/modules/mhsbiodataapi/builder"
	"tracerstudy-siak-service/pb"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func InitGrpc(server *grpc.Server, cfg config.Config, db *gorm.DB, grpcConn *grpc.ClientConn) {
	mhsbiodata := builder.BuildMhsBiodataApiHandler(cfg, grpcConn)
	pb.RegisterMhsBiodataApiServiceServer(server, mhsbiodata)
}
