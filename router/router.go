package router

import (
	"first-crud/controller"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(postController *controller.PostController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Print(w, "Welcome")
	})

	router.POST("/api/post", postController.Create)
	router.GET("/api/ ", postController.FindAll)
	router.GET("/api/post/:postId", postController.FindById)
	router.PUT("/api/post/:postId", postController.Update)
	router.DELETE("/api/post/:postId", postController.Delete)

	return router
}
