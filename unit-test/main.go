package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "test/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func main() {
	// hardcoded
	conn, err := grpc.Dial(":6666", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewPoliceReportingServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	responseChan := make(chan string)

	for i := 0; i < 70; i++ {
		go func() {
			_, err = client.FetchPolice(ctx, &pb.FetchPoliceRequest{
				UserInfo: &pb.GetPoliceUserEntry{
					UserLong:  -75.0364,
					UserLat:   36.8950,
					ZoomIndex: 100,
					City:      "chisinau",
				},
			})
			st, ok := status.FromError(err)

			if ok && st.Code() == codes.ResourceExhausted {
				responseChan <- "Reached the requests rate limit "
				return
			}
			if ok && st.Code() == codes.DeadlineExceeded {
				responseChan <- "Deadline exceeded"
				return
			}

		}()
	}

	response := <-responseChan
	fmt.Println("Result:", response)

}
