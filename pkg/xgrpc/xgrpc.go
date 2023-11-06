package xgrpc

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"runtime/debug"
)

// CrashInterceptor 恐慌捕获恢复服务端一元拦截器
func CrashInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	defer func() {
		if cause := recover(); cause != nil {
			fmt.Println(cause)
			debug.PrintStack()
			err = toPanicError(ctx, cause)
		}
	}()
	return handler(ctx, req)
}

// CrashStreamInterceptor 恐慌捕获恢复服务端流拦截器
func CrashStreamInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) (err error) {
	defer func() {
		if cause := recover(); cause != nil {
			ctx := context.Background()
			if ss != nil {
				ctx = ss.Context()
			}
			err = toPanicError(ctx, cause)
		}
	}()
	return handler(srv, ss)
}

// toPanicError 恐慌错误转换
func toPanicError(ctx context.Context, cause interface{}) error {
	logx.WithContext(ctx).Errorf("%+v [running]:\n%s", cause)
	return status.Errorf(codes.Internal, "panic: %v", cause)
}
