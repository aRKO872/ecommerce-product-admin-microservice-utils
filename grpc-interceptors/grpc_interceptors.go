package grpcinterceptors

import (
	"context"

	"google.golang.org/grpc"
)

func ErrorInterceptor() grpc.UnaryServerInterceptor {
	return  func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		resp, err = handler(ctx, req)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
}