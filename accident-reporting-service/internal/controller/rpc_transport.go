package controller

import (
	"context"
	"log"
	"net"
	"senpainikolay/pad-sem7/accident-reporting-service/internal/models"
	pb "senpainikolay/pad-sem7/accident-reporting-service/internal/pb"

	"golang.org/x/time/rate"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type IAccidentReportingService interface {
	Fetch(models.UserGeoInfo) (models.AccidentGeoInfoResponse, error)
	PostAccident(models.AccidentPostInfo) error
	ConfirmAccident(models.ConfirmationAccidentInfo) error
	CheckHealth(*rate.Limiter) models.HealthInfo
}

type AccidentReportingServer struct {
	pb.UnimplementedAccidentReportingServiceServer
	accReportingSvc IAccidentReportingService
}

var LIMITER = rate.NewLimiter(1, 1)

func Serve(acc_service IAccidentReportingService, bind string) {
	listener, err := net.Listen("tcp", bind)
	if err != nil {
		log.Fatalf("gRPC server error: failure to bind %v\n", bind)
	}

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(
			func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
				switch info.FullMethod {
				case "/proto.AccidentReportingService/HealthCheck":
					return handler(ctx, req)
				default:
					if !LIMITER.Allow() {
						return nil, status.Error(codes.ResourceExhausted, "rate limit exceeded")
					}
					return handler(ctx, req)
				}
			},
		),
	)

	accServer := AccidentReportingServer{accReportingSvc: acc_service}

	pb.RegisterAccidentReportingServiceServer(grpcServer, &accServer)
	log.Printf("gRPC API server listening on %v\n", bind)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("gRPC server error: %v\n", err)
	}
}

func (s *AccidentReportingServer) HealthCheck(ctx context.Context, req *pb.HealthRequest) (*pb.HealthResponse, error) {

	healthInfo := s.accReportingSvc.CheckHealth(LIMITER)

	return transformHealthMtoPB(healthInfo), nil

}

func transformHealthMtoPB(m models.HealthInfo) *pb.HealthResponse {
	return &pb.HealthResponse{
		Ready:    m.Ready,
		Database: m.Database,
		Load:     m.Load,
	}
}

func (s *AccidentReportingServer) PostAccident(ctx context.Context, req *pb.PostAccidentRequest) (*pb.GenericResponse, error) {

	err := s.accReportingSvc.PostAccident(transformPBpostToM(req.AccidentInfo))
	if err != nil {
		return &pb.GenericResponse{
			Error: true,
			Msg:   err.Error(),
		}, err
	}

	return &pb.GenericResponse{
		Error: false,
		Msg:   "added the accident succesfully",
	}, nil
}

func transformPBpostToM(pbM *pb.PostAccidentEntry) models.AccidentPostInfo {
	return models.AccidentPostInfo{
		Long:        pbM.AccidentLong,
		Lat:         pbM.AccidentLat,
		CarInvolved: pbM.CarsInvolved,
		StreetName:  pbM.StreetName,
		City:        pbM.City,
	}
}

func (s *AccidentReportingServer) FetchAccidents(ctx context.Context, req *pb.FetchAccidentRequest) (*pb.GetAccidentResponse, error) {

	accidentsData, err := s.accReportingSvc.Fetch(transformPBfetchToM(req.UserInfo))
	if err != nil {
		return &pb.GetAccidentResponse{}, err
	}
	return transformMfetchToPB(accidentsData), nil
}

func transformPBfetchToM(pbM *pb.GetAccidentUserEntry) models.UserGeoInfo {
	return models.UserGeoInfo{
		UserLong:  pbM.UserLong,
		UserLat:   pbM.UserLat,
		City:      pbM.City,
		ZoomIndex: pbM.ZoomIndex,
	}
}
func transformMfetchToPB(accData models.AccidentGeoInfoResponse) *pb.GetAccidentResponse {
	return &pb.GetAccidentResponse{
		AccidentInfo: transformMarrToPB(accData.Data),
	}
}

func transformMarrToPB(arr []models.AccidentInfo) []*pb.AccidentInfo {

	newArr := []*pb.AccidentInfo{}
	for _, val := range arr {
		newArr = append(newArr, &pb.AccidentInfo{
			AccidentLong:                     val.Long,
			AccidentLat:                      val.Lat,
			ConfirmationAccidentNotification: val.ConfirmationAccidentNotification,
			ConfirmationPoliceNotification:   val.ConfirmationPoliceNotification,
			ConfirmedBy:                      int64(val.ConfirmedBy),
		})
	}
	return newArr

}

func (s *AccidentReportingServer) ConfirmAccident(ctx context.Context, req *pb.ConfirmAccidentRequest) (*pb.GenericResponse, error) {

	err := s.accReportingSvc.ConfirmAccident(transformPBconfirmationToM(req.Info))
	if err != nil {
		return &pb.GenericResponse{
			Error: true,
			Msg:   err.Error(),
		}, err
	}
	return &pb.GenericResponse{
		Error: false,
		Msg:   "confirmation received",
	}, nil
}

func transformPBconfirmationToM(pbM *pb.ConfirmAccidentEntry) models.ConfirmationAccidentInfo {
	return models.ConfirmationAccidentInfo{
		Long:                 pbM.AccidentLong,
		Lat:                  pbM.AccidentLat,
		PoliceConfirmation:   pbM.PoliceConfirmation,
		AccidentConfirmation: pbM.AccidentConfirmation,
	}
}
