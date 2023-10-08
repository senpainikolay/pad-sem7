package controller

import (
	"context"
	"log"
	"net"
	"senpainikolay/pad-sem7/police-reporting-service/internal/models"
	pb "senpainikolay/pad-sem7/police-reporting-service/internal/pb"

	"google.golang.org/grpc"
)

type IPoliceReportingService interface {
	Fetch(models.UserGeoInfo) (models.PoliceGeoInfoResponse, error)
	PostPolice(models.PolicePostInfo) error
	ConfirmPolice(models.ConfirmationPoliceInfo) error
}

type PoliceReportingServer struct {
	pb.UnimplementedPoliceReportingServiceServer
	policeReportingSvc IPoliceReportingService
}

func Serve(police_service IPoliceReportingService, bind string) {
	listener, err := net.Listen("tcp", bind)
	if err != nil {
		log.Fatalf("gRPC server error: failure to bind %v\n", bind)
	}

	grpcServer := grpc.NewServer()

	policeServer := PoliceReportingServer{policeReportingSvc: police_service}

	pb.RegisterPoliceReportingServiceServer(grpcServer, &policeServer)
	log.Printf("gRPC API server listening on %v\n", bind)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("gRPC server error: %v\n", err)
	}
}

func (s *PoliceReportingServer) FetchPolice(ctx context.Context, req *pb.FetchPoliceRequest) (*pb.GetPoliceResponse, error) {

	policeData, err := s.policeReportingSvc.Fetch(transformPBfetchToM(req.UserInfo))
	if err != nil {
		return &pb.GetPoliceResponse{}, err
	}
	return transformMfetchToPB(policeData), nil
}

func transformPBfetchToM(pbM *pb.GetPoliceUserEntry) models.UserGeoInfo {
	return models.UserGeoInfo{
		UserLong:  pbM.UserLong,
		UserLat:   pbM.UserLat,
		City:      pbM.City,
		ZoomIndex: pbM.ZoomIndex,
	}
}
func transformMfetchToPB(policeData models.PoliceGeoInfoResponse) *pb.GetPoliceResponse {
	return &pb.GetPoliceResponse{
		PoliceInfo: transformMarrToPB(policeData.Data),
	}
}

func transformMarrToPB(arr []models.PoliceInfo) []*pb.PoliceInfo {

	newArr := []*pb.PoliceInfo{}
	for _, val := range arr {
		newArr = append(newArr, &pb.PoliceInfo{
			PolLong:                  val.PolLong,
			PolLat:                   val.PolLat,
			ConfirmationNotification: val.ConfirmationNotification,
			ConfirmedBy:              int64(val.ConfirmedBy),
		})
	}
	return newArr

}

func (s *PoliceReportingServer) PostPolice(ctx context.Context, req *pb.PostPoliceRequest) (*pb.PoliceResponse, error) {

	err := s.policeReportingSvc.PostPolice(transformPBpostToM(req.PoliceInfo))
	if err != nil {
		return &pb.PoliceResponse{
			Error: true,
			Msg:   err.Error(),
		}, err
	}

	return &pb.PoliceResponse{
		Error: false,
		Msg:   "added the police succesfully",
	}, nil
}

func transformPBpostToM(pbM *pb.PostPoliceEntry) models.PolicePostInfo {
	return models.PolicePostInfo{
		PolLong: pbM.PolLong,
		PolLat:  pbM.PolLat,
		City:    pbM.City,
	}
}

func (s *PoliceReportingServer) ConfirmPolice(ctx context.Context, req *pb.ConfirmPoliceRequest) (*pb.PoliceResponse, error) {

	err := s.policeReportingSvc.ConfirmPolice(transformPBconfirmationToM(req.PoliceInfo))
	if err != nil {
		return &pb.PoliceResponse{
			Error: true,
			Msg:   err.Error(),
		}, err
	}

	return &pb.PoliceResponse{
		Error: false,
		Msg:   "confirmation received",
	}, nil
}

func transformPBconfirmationToM(pbM *pb.ConfirmPoliceEntry) models.ConfirmationPoliceInfo {
	return models.ConfirmationPoliceInfo{
		PoliceInfo: models.PolicePostInfo{
			PolLong: pbM.PolLong,
			PolLat:  pbM.PolLat,
			City:    pbM.City,
		},
		Confirmation: pbM.Confirmation,
	}
}
