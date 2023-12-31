package service

import (
	"context"
	"first-crud/data/request"
	"first-crud/data/response"
)

type PostService interface {
	Create(ctx context.Context, request request.PostCreateRequest)
	Update(ctx context.Context, request request.PostUpdateRequest)
	Delete(ctx context.Context, postId string)
	FindById(ctx context.Context, postId string) response.PostResponse
	FindAll(ctx context.Context) []response.PostResponse
}
