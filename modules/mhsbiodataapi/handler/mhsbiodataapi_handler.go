package handler

import (
	"context"
	"log"
	"net/http"
	"tracerstudy-siak-service/common/config"
	"tracerstudy-siak-service/common/errors"
	"tracerstudy-siak-service/modules/mhsbiodataapi/entity"
	"tracerstudy-siak-service/modules/mhsbiodataapi/service"
	"tracerstudy-siak-service/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type MhsBiodataApiHandler struct {
	pb.UnimplementedMhsBiodataApiServiceServer
	config        config.Config
	mhsbiodataSvc service.MhsBiodataApiServiceUseCase
}

func NewMhsBiodataHandler(config config.Config, mhsbiodataApiService service.MhsBiodataApiServiceUseCase) *MhsBiodataApiHandler {
	return &MhsBiodataApiHandler{
		config:        config,
		mhsbiodataSvc: mhsbiodataApiService,
	}
}

func (mbh *MhsBiodataApiHandler) FetchMhsBiodataByNim(ctx context.Context, req *pb.MhsBiodataApiRequest) (*pb.MhsBiodataApiResponse, error) {
	nim := req.GetNim()

	var apiResponse *entity.MhsBiodataApi
	apiResponse, err := mbh.mhsbiodataSvc.FetchMhsBiodataByNimFromSiakApi(nim)
	if err != nil {
		if apiResponse == nil {
			log.Println("WARNING: [MhsBiodataHandler - FetchMhsBiodataByNim] Resource not found: nim ", nim)
			// return nil, status.Errorf(codes.NotFound, "resource not found")
			return &pb.MhsBiodataApiResponse{
				Code:    uint32(http.StatusNotFound),
				Message: "mhsbiodata not found",
			}, status.Errorf(codes.NotFound, "mhsbiodata not found")
		}

		parseError := errors.ParseError(err)
		log.Println("ERROR: [MhsBiodataHandler - FetchMhsBiodataByNim] Error while fetching mhs biodata:", parseError.Message)
		// return nil, status.Errorf(parseError.Code, parseError.Message)
		return &pb.MhsBiodataApiResponse{
			Code:    uint32(http.StatusInternalServerError),
			Message: parseError.Message,
		}, status.Errorf(parseError.Code, parseError.Message)
	}

	var mhsBiodata = entity.ConvertEntityToProto(apiResponse)

	return &pb.MhsBiodataApiResponse{
		Code:    uint32(http.StatusOK),
		Message: "get mhs biodata success",
		Data:    mhsBiodata,
	}, nil
}

func (mbh *MhsBiodataApiHandler) CheckMhsAlumni(ctx context.Context, req *pb.CheckMhsAlumniRequest) (*pb.CheckMhsAlumniResponse, error) {
	nim := req.GetNim()
	tglSidang := req.GetTglSidang()

	var apiResponse *entity.MhsBiodataApi
	apiResponse, err := mbh.mhsbiodataSvc.FetchMhsBiodataByNimFromSiakApi(nim)
	if err != nil {
		if apiResponse == nil {
			log.Println("WARNING: [MhsBiodataHandler - FetchMhsBiodataByNim] Resource not found: nim ", nim)
			// return nil, status.Errorf(codes.NotFound, "resource not found")
			return &pb.CheckMhsAlumniResponse{
				Code:     uint32(http.StatusNotFound),
				Message:  "mahasiswa not found",
				IsAlumni: false,
			}, status.Errorf(codes.NotFound, "mahasiswa not found")
		}

		parseError := errors.ParseError(err)
		log.Println("ERROR: [MhsBiodataHandler - FetchMhsBiodataByNim] Error while fetching mhs biodata:", parseError.Message)
		// return nil, status.Errorf(parseError.Code, parseError.Message)
		return &pb.CheckMhsAlumniResponse{
			Code:     uint32(http.StatusInternalServerError),
			Message:  parseError.Message,
			IsAlumni: false,
		}, status.Errorf(parseError.Code, parseError.Message)
	}

	if apiResponse.TGLSIDANG != tglSidang {
		log.Println("WARNING: [MhsBiodataHandler - CheckMhsAlumni] Tanggal sidang not match")
		return &pb.CheckMhsAlumniResponse{
			Code:     uint32(http.StatusOK),
			Message:  "tanggal sidang not match",
			IsAlumni: false,
		}, nil
	}

	if apiResponse.KODESTATUS == "2" {
		return &pb.CheckMhsAlumniResponse{
			Code:     uint32(http.StatusOK),
			Message:  "get mhs status alumni success",
			IsAlumni: true,
		}, nil
	} else {
		log.Println("WARNING: [MhsBiodataHandler - CheckMhsAlumni] Mahasiswa is not alumni yet")
		return &pb.CheckMhsAlumniResponse{
			Code:     uint32(http.StatusOK),
			Message:  "mahasiswa is not alumni yet",
			IsAlumni: false,
		}, nil
	}
}
