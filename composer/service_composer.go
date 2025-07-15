package composer

import (
	"blog-tech/common"
	categorybiz "blog-tech/internal/categories/business"
	categorypb "blog-tech/internal/categories/proto/pb"
	categorymysql "blog-tech/internal/categories/repository/mysql"
	categorystorerpc "blog-tech/internal/categories/repository/rpc"
	categoryapi "blog-tech/internal/categories/transport/api"
	categoryrpc "blog-tech/internal/categories/transport/rpc"
	userbiz "blog-tech/internal/users/business"
	userpb "blog-tech/internal/users/proto/pb"
	usermysql "blog-tech/internal/users/repository/mysql"
	sctx "blog-tech/plugin"
	"log"
	"os"

	"github.com/gin-gonic/gin"

	userAPI "blog-tech/internal/users/transport/api"
	userrpc "blog-tech/internal/users/transport/rpc"
)

type UserService interface {
	RegisterHdl() gin.HandlerFunc
	LoginHdl() gin.HandlerFunc
}

type CategoryService interface {
	CreateCategoryHdl() gin.HandlerFunc
	UpdateCategoryHdl() gin.HandlerFunc
	GetCategoryByIDHdl() gin.HandlerFunc
}

func ComposeUserService(serviceContext sctx.ServiceContext) UserService {
	db := serviceContext.MustGet(common.KeyCompMySQL).(common.GormComponent)
	userRepo := usermysql.NewUserRepository(db.GetDB())

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Fatal("JWT_SECRET is not set")
	}

	jwtManager := common.NewJwtManager(secret)

	biz := userbiz.NewUserBusiness(userRepo, jwtManager)
	serviceAPI := userAPI.NewHandler(biz)

	return serviceAPI
}

func ComposeUserGRPCService(serviceCtx sctx.ServiceContext) userpb.UserServiceServer {
	db := serviceCtx.MustGet(common.KeyCompMySQL).(common.GormComponent)
	// jwtComp := serviceCtx.MustGet(common.KeyCompJWT).(common.JwtManager)

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Fatal("JWT_SECRET is not set")
	}

	jwtManager := common.NewJwtManager(secret)

	userRepo := usermysql.NewUserRepository(db.GetDB())

	biz := userbiz.NewUserBusiness(userRepo, jwtManager)
	authService := userrpc.NewUserService(biz)

	return authService
}

func ComposeCategoryService(serviceContext sctx.ServiceContext) CategoryService {
	db := serviceContext.MustGet(common.KeyCompMySQL).(common.GormComponent)
	categoryRepo := categorymysql.NewCategoryStore(db.GetDB())

	userClient := categorystorerpc.NewClient(ComposeUserRPCClient(serviceContext))

	biz := categorybiz.NewCategoryBusiness(categoryRepo, userClient)
	serviceAPI := categoryapi.NewApi(biz)

	return serviceAPI
}

func ComposeCategoryGRPCService(serviceCtx sctx.ServiceContext) categorypb.CategoryServiceServer {
	db := serviceCtx.MustGet(common.KeyCompMySQL).(common.GormComponent)
	categoryRepo := categorymysql.NewCategoryStore(db.GetDB())

	biz := categorybiz.NewCategoryBusiness(categoryRepo, nil)

	categoryService := categoryrpc.NewCategoryService(biz)

	return categoryService
}
