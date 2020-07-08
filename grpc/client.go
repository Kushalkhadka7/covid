package grpcconn

import (
	"fmt"

	"google.golang.org/grpc"
)

//GrpcClient connection.
type GrpcClient struct {
	*grpc.ClientConn
}

// NewGrpc grpc client to connect to the server in given port number.
func NewGrpc(port int) (*GrpcClient, error) {

	conn, err := grpc.Dial(fmt.Sprintf(":%d", port), grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("Cannot connect to grpc server:%v", err)

	}

	fmt.Printf("Grpc Server listening on :%d\n", port)

	return &GrpcClient{conn}, nil
}
