package client

import (
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "demo/proto"
	"demo/utils/grpchelper/config"
)

func DemoClient(addr string) (pb.DemoServiceClient, error) {
	config, err := config.GenerateServiceConfig(
		"round_robin",
		[]string{pb.DemoService_ServiceDesc.ServiceName},
		[]string{},
		5,
		1*time.Second,
		10*time.Second,
		2.0,
		[]string{"UNAVAILABLE"},
	)
	if err != nil {
		return nil, err
	}

	// config := `{
	//     "loadBalancingPolicy": "round_robin",
	//     "methodConfig": [
	//         {
	//             "name": [{"service": "` + pb.DemoService_ServiceDesc.ServiceName + `"}],
	//             "waitForReady": true,
	//             "timeout": "1s",
	//             "retryPolicy": {
	//                 "maxAttempts": 5,
	//                 "initialBackoff": "1s",
	//                 "maxBackoff": "10s",
	//                 "backoffMultiplier": 2.0,
	//                 "retryableStatusCodes": ["UNAVAILABLE"]
	//             }
	//         }
	//     ]
	// }`

	conn, err := grpc.NewClient(
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(config),
	)
	if err != nil {
		return nil, err
	}

	return pb.NewDemoServiceClient(conn), nil
}
