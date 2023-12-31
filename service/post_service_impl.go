package service

import (
	"context"
	"first-crud/data/request"
	"first-crud/data/response"
	"first-crud/helper"
	"first-crud/model"
	"first-crud/repository"
)

type PostServiceImpl struct {
	PostRepository repository.PostRepository
}

func NewPostServiceImpl(postRepository repository.PostRepository) PostService {
	return &PostServiceImpl{
		PostRepository: postRepository,
	}
}

// Create implements PostService.
func (p *PostServiceImpl) Create(ctx context.Context, request request.PostCreateRequest) {
	postData := model.Post{
		Title:       request.Title,
		Published:   request.Published,
		Description: request.Description,
	}
	p.PostRepository.Save(ctx, postData)
}

// Delete implements PostService.
func (p *PostServiceImpl) Delete(ctx context.Context, postId string) {
	post, err := p.PostRepository.FindById(ctx, postId)
	helper.ErrorPanic(err)
	p.PostRepository.Delete(ctx, post.Id)
}

// FindAll implements PostService.
func (p *PostServiceImpl) FindAll(ctx context.Context) []response.PostResponse {
	posts := p.PostRepository.FindAll(ctx)

	var postResp []response.PostResponse

	for _, value := range posts {
		post := response.PostResponse{
			Id:          value.Id,
			Title:       value.Title,
			Published:   value.Published,
			Description: value.Description,
		}
		postResp = append(postResp, post)
	}

	return postResp
}

// FindById implements PostService.
func (p *PostServiceImpl) FindById(ctx context.Context, postId string) response.PostResponse {
	post, err := p.PostRepository.FindById(ctx, postId)
	helper.ErrorPanic(err)

	postResponse := response.PostResponse{
		Id:          post.Id,
		Title:       post.Title,
		Published:   post.Published,
		Description: post.Description,
	}

	return postResponse
}

// Update implements PostService.
func (p *PostServiceImpl) Update(ctx context.Context, request request.PostUpdateRequest) {
	postData := model.Post{
		Id:          request.Id,
		Title:       request.Title,
		Published:   request.Published,
		Description: request.Description,
	}
	p.PostRepository.Update(ctx, postData)
}
