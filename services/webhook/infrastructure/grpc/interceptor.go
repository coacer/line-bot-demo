package grpc

import (
	"context"
	"log"

	"google.golang.org/grpc"
)

type Interceptor struct{}

func (i *Interceptor) logging() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		log.Printf("before [%s] is called with req [%+v]", info.FullMethod, req)
		resp, err = handler(ctx, req)
		if err != nil {
			log.Printf("error occured in method [%s]: [%+v]", info.FullMethod, err)
			return nil, err
		}
		log.Printf("after [%s] is called: response=[%+v]", info.FullMethod, resp)
		return
	}
}
