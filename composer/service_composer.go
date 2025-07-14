package composer

import (
	"blog-tech/common"
	userbiz "blog-tech/internal/users/biz"
	"blog-tech/internal/users/proto/pb"
	usermysql "blog-tech/internal/users/repository/mysql"
	sctx "blog-tech/plugin"
	"log"
	"os"

	"github.com/gin-gonic/gin"

	userAPI "blog-tech/internal/users/transport/api"
	"blog-tech/internal/users/transport/rpc"
)

type UserService interface {
	Register() gin.HandlerFunc
	Login() gin.HandlerFunc
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

func ComposeUserGRPCService(serviceCtx sctx.ServiceContext) pb.UserServiceServer {
	db := serviceCtx.MustGet(common.KeyCompMySQL).(common.GormComponent)
	// jwtComp := serviceCtx.MustGet(common.KeyCompJWT).(common.JwtManager)

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Fatal("JWT_SECRET is not set")
	}

	jwtManager := common.NewJwtManager(secret)

	userRepo := usermysql.NewUserRepository(db.GetDB())

	biz := userbiz.NewUserBusiness(userRepo, jwtManager)
	authService := rpc.NewService(biz)

	return authService
}
