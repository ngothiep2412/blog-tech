package composer

import (
	"blog-tech/common"
	categorypb "blog-tech/internal/categories/proto/pb"
	userpb "blog-tech/internal/users/proto/pb"
	sctx "blog-tech/plugin"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ComposeUserRPCClient(serviceCtx sctx.ServiceContext) userpb.UserServiceClient {
	configComp := serviceCtx.MustGet(common.KeyCompConf).(common.ConfigComponent)

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.NewClient(configComp.GetGRPCUserAddress(), opts...)
	if err != nil {
		log.Fatal(err)
	}

	return userpb.NewUserServiceClient(conn)
}

func ComposeCategoryRPCClient(serviceCtx sctx.ServiceContext) categorypb.CategoryServiceClient {
	configComp := serviceCtx.MustGet(common.KeyCompConf).(common.ConfigComponent)

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.NewClient(configComp.GetGRPCCategoryAddress(), opts...)

	if err != nil {
		log.Fatal(err)
	}

	return categorypb.NewCategoryServiceClient(conn)
}
