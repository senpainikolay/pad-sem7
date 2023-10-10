package policeserviceclient

import (
	"context"
	"log"
	"time"

	"senpainikolay/pad-sem7/accident-reporting-service/internal/models"
	pb "senpainikolay/pad-sem7/accident-reporting-service/internal/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InformExternalService(data models.ExernalServiceData, long, lat float64) error {
	conn, err := grpc.Dial(":6666", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewPoliceReportingServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	res, err := client.FetchPolice(ctx, &pb.FetchPoliceRequest{
		UserInfo: &pb.GetPoliceUserEntry{
			UserLong:  long,
			UserLat:   lat,
			ZoomIndex: 100,
			City:      data.City,
		},
	})
	if err != nil {
		return nil
	}
	for _, v := range res.PoliceInfo {
		data.NearbyPolice = append(data.NearbyPolice,
			models.PoliceInfo{
				Long:        v.PolLong,
				Lat:         v.PolLat,
				ConfirmedBy: int(v.ConfirmedBy),
			},
		)

	}

	log.Println("OVERALL DATA SETN TO EXTERNAL SERVICE:::")
	log.Println(data)

	return nil

}
