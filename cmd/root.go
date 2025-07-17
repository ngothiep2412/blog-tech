package cmd

import (
	"blog-tech/common"
	"blog-tech/composer"
	articletagpb "blog-tech/internal/article_tags/proto/pb"
	categorypb "blog-tech/internal/categories/proto/pb"
	tagpb "blog-tech/internal/tags/proto/pb"
	userpb "blog-tech/internal/users/proto/pb"
	"blog-tech/middleware"
	sctx "blog-tech/plugin"
	"blog-tech/plugin/ginc"
	"blog-tech/plugin/gormc"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

func NewServiceCtx() sctx.ServiceContext {
	return sctx.NewServiceContext(
		sctx.WithName("Blog Tech Service"),
		sctx.WithComponent(ginc.NewGin(common.KeyCompGIN)),
		sctx.WithComponent(gormc.NewGormDB(common.KeyCompMySQL, "")),
		sctx.WithComponent(NewConfig()),
	)
}

var rootCmd = &cobra.Command{
	Use:   "blog-tech",
	Short: "Blog Tech",
	Long:  "A tech blog application",
	Run: func(cmd *cobra.Command, args []string) {
		serviceCtx := NewServiceCtx()

		logger := sctx.GlobalLogger().GetLogger("service")

		time.Sleep(time.Second * 5)

		if err := serviceCtx.Load(); err != nil {
			logger.Fatal(err)
		}

		ginComp := serviceCtx.MustGet(common.KeyCompGIN).(common.GINComponent)

		router := ginComp.GetRouter()
		router.Use(gin.Recovery(), gin.Logger())

		router.Use()

		router.SetTrustedProxies([]string{"127.0.0.1"})

		go StartGRPCServices(serviceCtx)

		v1 := router.Group("/v1")

		SetupRoutes(v1, serviceCtx)

		if err := router.Run(fmt.Sprintf(":%d", ginComp.GetPort())); err != nil {
			logger.Fatal(err)
		}

	},
}

func SetupRoutes(router *gin.RouterGroup, serviceCtx sctx.ServiceContext) {
	userAPIService := composer.ComposeUserService(serviceCtx)
	categoryAPIService := composer.ComposeCategoryService(serviceCtx)
	tagAPIService := composer.ComposeTagService(serviceCtx)
	articleAPIService := composer.ComposeArticleService(serviceCtx)
	articleLikeAPIService := composer.ComposeArticleLikeService(serviceCtx)

	router.POST("/register", userAPIService.RegisterHdl())
	router.POST("/login", userAPIService.LoginHdl())
	router.POST("/refresh-token", userAPIService.RefreshTokenHdl())

	router.POST("/categories", middleware.RequireAuth(), categoryAPIService.CreateCategoryHdl())
	router.PUT("/categories/:id", middleware.RequireAuth(), categoryAPIService.UpdateCategoryHdl())
	router.GET("/categories/:id", middleware.RequireAuth(), categoryAPIService.GetCategoryByIDHdl())

	router.POST("/tags", middleware.RequireAuth(), tagAPIService.CreateTagHdl())
	router.PUT("/tags/:id", middleware.RequireAuth(), tagAPIService.UpdateTagHdl())
	router.GET("/tags/:id", middleware.RequireAuth(), tagAPIService.GetTagByIDHdl())

	router.POST("/articles", middleware.RequireAuth(), articleAPIService.CreateArticleHdl())

	router.POST("/articles/:article_id/like", middleware.RequireAuth(), articleLikeAPIService.LikeArticleHdl())
	router.DELETE("/articles/:article_id/like", middleware.RequireAuth(), articleLikeAPIService.UnlikeArticleHdl())
}

func StartGRPCServices(serviceCtx sctx.ServiceContext) {
	configComp := serviceCtx.MustGet(common.KeyCompConf).(common.ConfigComponent)

	logger := serviceCtx.Logger("grpc")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", configComp.GetGRPCPort()))

	if err != nil {
		log.Fatal(err)
	}

	logger.Infof("GRPC Server is listening on %d ...\n", configComp.GetGRPCPort())
	logger.Infof("GRPC User Server is listening on %s ...\n", configComp.GetGRPCUserAddress())
	logger.Infof("GRPC Category Server is listening on %s ...\n", configComp.GetGRPCCategoryAddress())
	logger.Infof("GRPC Tag Server is listening on %s ...\n", configComp.GetGRPCTagAddress())
	logger.Infof("GRPC Article Tag Server is listening on %s ...\n", configComp.GetGRPCArticleTagAddress())

	s := grpc.NewServer()

	userpb.RegisterUserServiceServer(s, composer.ComposeUserGRPCService(serviceCtx))
	categorypb.RegisterCategoryServiceServer(s, composer.ComposeCategoryGRPCService(serviceCtx))
	tagpb.RegisterTagServiceServer(s, composer.ComposeTagGRPCService(serviceCtx))
	articletagpb.RegisterArticleTagServiceServer(s, composer.ComposeArticleTagGRPCService(serviceCtx))

	if err := s.Serve(lis); err != nil {
		log.Fatalln(err)
	}
}

func Execute() {
	rootCmd.AddCommand(outenvCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
