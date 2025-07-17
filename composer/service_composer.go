package composer

import (
	"blog-tech/common"
	articletagbiz "blog-tech/internal/article_tags/business"
	articletagpb "blog-tech/internal/article_tags/proto/pb"
	articletagreopomysql "blog-tech/internal/article_tags/repository/mysql"
	articletagrpc "blog-tech/internal/article_tags/transport/rpc"
	articlebiz "blog-tech/internal/articles/business"
	articlerepomysql "blog-tech/internal/articles/repository/mysql"
	articlereporpc "blog-tech/internal/articles/repository/rpc"
	articleapi "blog-tech/internal/articles/transport/api"
	categorybiz "blog-tech/internal/categories/business"
	categorypb "blog-tech/internal/categories/proto/pb"
	categorymysql "blog-tech/internal/categories/repository/mysql"
	categorystorerpc "blog-tech/internal/categories/repository/rpc"
	categoryapi "blog-tech/internal/categories/transport/api"
	categoryrpc "blog-tech/internal/categories/transport/rpc"
	tagbiz "blog-tech/internal/tags/business"
	tagpb "blog-tech/internal/tags/proto/pb"
	tagrepomysql "blog-tech/internal/tags/repository/mysql"
	tagreporpc "blog-tech/internal/tags/repository/rpc"
	tagapi "blog-tech/internal/tags/transport/api"
	tagrpc "blog-tech/internal/tags/transport/rpc"
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
	RefreshTokenHdl() gin.HandlerFunc
}

type CategoryService interface {
	CreateCategoryHdl() gin.HandlerFunc
	UpdateCategoryHdl() gin.HandlerFunc
	GetCategoryByIDHdl() gin.HandlerFunc
}

type TagService interface {
	CreateTagHdl() gin.HandlerFunc
	GetTagByIDHdl() gin.HandlerFunc
	UpdateTagHdl() gin.HandlerFunc
}

type ArticleService interface {
	CreateArticleHdl() gin.HandlerFunc
}

func ComposeUserService(serviceContext sctx.ServiceContext) UserService {
	db := serviceContext.MustGet(common.KeyCompMySQL).(common.GormComponent)
	userRepo := usermysql.NewUserRepository(db.GetDB())

	secret := os.Getenv("JWT_SECRET")
	refreshSecret := os.Getenv("JWT_REFRESH_SECRET")

	if secret == "" {
		log.Fatal("JWT_SECRET is not set")
	}
	if refreshSecret == "" {
		log.Fatal("JWT_REFRESH_SECRET is not set")
	}

	jwtManager := common.NewJwtManager(secret, refreshSecret)

	biz := userbiz.NewUserBusiness(userRepo, jwtManager)
	serviceAPI := userAPI.NewHandler(biz)

	return serviceAPI
}

func ComposeUserGRPCService(serviceCtx sctx.ServiceContext) userpb.UserServiceServer {
	db := serviceCtx.MustGet(common.KeyCompMySQL).(common.GormComponent)
	// jwtComp := serviceCtx.MustGet(common.KeyCompJWT).(common.JwtManager)

	secret := os.Getenv("JWT_SECRET")
	refreshSecret := os.Getenv("JWT_REFRESH_SECRET")
	if secret == "" {
		log.Fatal("JWT_SECRET is not set")
	}
	if refreshSecret == "" {
		log.Fatal("JWT_REFRESH_SECRET is not set")
	}

	jwtManager := common.NewJwtManager(secret, refreshSecret)

	userRepo := usermysql.NewUserRepository(db.GetDB())

	biz := userbiz.NewUserBusiness(userRepo, jwtManager)
	authService := userrpc.NewUserService(biz)

	return authService
}

func ComposeCategoryService(serviceContext sctx.ServiceContext) CategoryService {
	db := serviceContext.MustGet(common.KeyCompMySQL).(common.GormComponent)
	categoryRepo := categorymysql.NewCategoryRepository(db.GetDB())

	userClient := categorystorerpc.NewClient(ComposeUserRPCClient(serviceContext))

	biz := categorybiz.NewCategoryBusiness(categoryRepo, userClient)
	serviceAPI := categoryapi.NewApi(biz)

	return serviceAPI
}

func ComposeCategoryGRPCService(serviceCtx sctx.ServiceContext) categorypb.CategoryServiceServer {
	db := serviceCtx.MustGet(common.KeyCompMySQL).(common.GormComponent)
	categoryRepo := categorymysql.NewCategoryRepository(db.GetDB())

	biz := categorybiz.NewCategoryBusiness(categoryRepo, nil)

	categoryService := categoryrpc.NewCategoryService(biz)

	return categoryService
}

func ComposeTagService(serviceContext sctx.ServiceContext) TagService {
	db := serviceContext.MustGet(common.KeyCompMySQL).(common.GormComponent)
	tagRepo := tagrepomysql.NewTagRepository(db.GetDB())

	userClient := tagreporpc.NewClient(ComposeUserRPCClient(serviceContext))

	biz := tagbiz.NewBusiness(tagRepo, userClient)

	serviceAPI := tagapi.NewApi(biz)

	return serviceAPI
}

func ComposeTagGRPCService(serviceCtx sctx.ServiceContext) tagpb.TagServiceServer {
	db := serviceCtx.MustGet(common.KeyCompMySQL).(common.GormComponent)
	tagRepo := tagrepomysql.NewTagRepository(db.GetDB())

	biz := tagbiz.NewBusiness(tagRepo, nil)

	tagService := tagrpc.NewTagService(biz)

	return tagService
}

// Article-Tag
func ComposeArticleTagGRPCService(serviceCtx sctx.ServiceContext) articletagpb.ArticleTagServiceServer {
	db := serviceCtx.MustGet(common.KeyCompMySQL).(common.GormComponent)
	articleTagRepo := articletagreopomysql.NewArticleTagRepository(db.GetDB())

	biz := articletagbiz.NewArticleTagBusiness(articleTagRepo)

	articleTagService := articletagrpc.NewArticleTagService(biz)

	return articleTagService
}

// Article
func ComposeArticleService(serviceContext sctx.ServiceContext) ArticleService {
	// Get database component
	db := serviceContext.MustGet(common.KeyCompMySQL).(common.GormComponent)

	// Initialize article repository
	articleRepo := articlerepomysql.NewArticleRepository(db.GetDB())

	// Create RPC client
	rpcClient := articlereporpc.NewClient(
		ComposeUserRPCClient(serviceContext),
		ComposeCategoryRPCClient(serviceContext),
		ComposeTagRPCClient(serviceContext),
		ComposeArticleTagRPCClient(serviceContext),
	)

	// Initialize business layer with repositories
	biz := articlebiz.NewArticleBusiness(
		articleRepo,
		rpcClient, // UserRepository
		rpcClient, // ArticleTagRepository
		rpcClient, // TagRepository
		rpcClient, // CategoryRepository
	)

	// Create API handler with business logic
	serviceAPI := articleapi.NewArticleApi(biz)

	return serviceAPI
}
