package composer

import (
	"blog-tech/common"
	articletagpb "blog-tech/internal/article_tags/proto/pb"
	categorypb "blog-tech/internal/categories/proto/pb"
	tagpb "blog-tech/internal/tags/proto/pb"
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

func ComposeTagRPCClient(serviceCtx sctx.ServiceContext) tagpb.TagServiceClient {
	configComp := serviceCtx.MustGet(common.KeyCompConf).(common.ConfigComponent)

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.NewClient(configComp.GetGRPCUserAddress(), opts...)
	if err != nil {
		log.Fatal(err)
	}

	return tagpb.NewTagServiceClient(conn)
}

func ComposeArticleTagRPCClient(serviceCtx sctx.ServiceContext) articletagpb.ArticleTagServiceClient {
	configComp := serviceCtx.MustGet(common.KeyCompConf).(common.ConfigComponent)

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.NewClient(configComp.GetGRPCUserAddress(), opts...)
	if err != nil {
		log.Fatal(err)
	}

	return articletagpb.NewArticleTagServiceClient(conn)
}
