package controller

import (
	"context"
	"log"
	"net"
	"senpainikolay/pad-sem7/accident-reporting-service/internal/models"
	pb "senpainikolay/pad-sem7/accident-reporting-service/internal/pb"

	"google.golang.org/grpc"
)

type IAccidentReportingService interface {
	Fetch(models.UserGeoInfo) (models.AccidentGeoInfoResponse, error)
	PostAccident(models.AccidentPostInfo) error
	//ConfirmPolice(models.ConfirmationAccidentInfo) error
}

type AccidentReportingServer struct {
	pb.UnimplementedAccidentReportingServiceServer
	accReportingSvc IAccidentReportingService
}

func Serve(acc_service IAccidentReportingService, bind string) {
	listener, err := net.Listen("tcp", bind)
	if err != nil {
		log.Fatalf("gRPC server error: failure to bind %v\n", bind)
	}

	grpcServer := grpc.NewServer()

	accServer := AccidentReportingServer{accReportingSvc: acc_service}

	pb.RegisterAccidentReportingServiceServer(grpcServer, &accServer)
	log.Printf("gRPC API server listening on %v\n", bind)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("gRPC server error: %v\n", err)
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
