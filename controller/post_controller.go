package controller

import (
	"first-crud/data/request"
	"first-crud/data/response"
	"first-crud/helper"
	"first-crud/service"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type PostController struct {
	PostService service.PostService
}

func NewPostController(postService service.PostService) *PostController {
	return &PostController{
		PostService: postService,
	}
}

// create function
func (controller *PostController) Create(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	postCreateRequest := request.PostCreateRequest{}
	helper.ReadRequestBody(requests, &postCreateRequest)

	controller.PostService.Create(requests.Context(), postCreateRequest)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   nil,
	}
	helper.WriteResponseBody(writer, webResponse)
}

// update function
func (controller *PostController) Update(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	PostUpdateRequest := request.PostUpdateRequest{}
	helper.ReadRequestBody(requests, &PostUpdateRequest)

	postId := params.ByName("postId")
	PostUpdateRequest.Id = postId

	controller.PostService.Update(requests.Context(), PostUpdateRequest)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   nil,
	}
	helper.WriteResponseBody(writer, webResponse)
}

// FinAll function
func (controller *PostController) FindAll(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	result := controller.PostService.FindAll(requests.Context())
	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   result,
	}
	helper.WriteResponseBody(writer, webResponse)
}

// find by id function
func (controller *PostController) FindById(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {

	postId := params.ByName("postId")
	result := controller.PostService.FindById(requests.Context(), postId)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   result,
	}
	helper.WriteResponseBody(writer, webResponse)
}

// Delete id function
func (controller *PostController) Delete(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {

	postId := params.ByName("postId")
	controller.PostService.Delete(requests.Context(), postId)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   nil,
	}
	helper.WriteResponseBody(writer, webResponse)
}
