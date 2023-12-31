package main

import (
	"first-crud/config"
	"first-crud/controller"
	"first-crud/helper"
	"first-crud/repository"
	"first-crud/router"
	"first-crud/service"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("🚀 Could not load environments variables", err)
	}
	fmt.Print("Start server! " + os.Getenv("PORT"+"\n"))

	//conexao com o banco de dados
	db, err := config.ConnectDb()
	helper.ErrorPanic(err)
	defer db.Prisma.Disconnect()

	//repositorio
	postRepository := repository.NewPostRepository(db)

	// service
	postService := service.NewPostServiceImpl(postRepository)

	//controller
	postController := controller.NewPostController(postService)

	//router
	routes := router.NewRouter(postController)

	server := &http.Server{
		Addr:           ":" + os.Getenv(("PORT")),
		Handler:        routes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	server_err := server.ListenAndServe()
	if server_err != nil {
		panic(server_err)
	}

}
