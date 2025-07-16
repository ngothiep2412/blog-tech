package articletagrpc

import (
	articletagpb "blog-tech/internal/article_tags/proto/pb"
	"context"
	"fmt"
)

type ArticleTagBusiness interface {
	CreateArticleTag(ctx context.Context, articleID int, tagIDs []int) error
}

type articleTagService struct {
	business ArticleTagBusiness
}

func NewArticleTagService(business ArticleTagBusiness) *articleTagService {
	return &articleTagService{
		business: business,
	}
}

func (s *articleTagService) CreateArticleTags(ctx context.Context, req *articletagpb.CreateArticleTagsRequest) (*articletagpb.CreateArticleTagsResponse, error) {
	tagIds := make([]int, len(req.TagIds))

	for i, tagID := range req.TagIds {
		tagIds[i] = int(tagID)
	}

	err := s.business.CreateArticleTag(ctx, int(req.ArticleId), tagIds)

	if err != nil {
		return &articletagpb.CreateArticleTagsResponse{
			Success: false,
			Message: fmt.Sprintf("Failed to create article tags: %v", err),
		}, nil
	}

	return &articletagpb.CreateArticleTagsResponse{
		Success: true,
		Message: "Create article tags successfully",
	}, nil
}
