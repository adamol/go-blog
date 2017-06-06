package main

import (
    "net/http"
    "log"

    "github.com/julienschmidt/httprouter"

    "blog/controllers"
)

func main() {
    router := httprouter.New()

    postsController := controllers.PostsController{}

    router.GET("/", postsController.Index)
    router.POST("/", postsController.Store)
    router.GET("/posts/:slug", postsController.Show)

    commentsController := controllers.CommentsController{}

    router.POST("/posts/:slug/comments", commentsController.Store)

    authController := controllers.AuthController{}

    router.GET("/login/form", authController.LoginForm)
    router.POST("/login", authController.HandleLogin)

    router.GET("/register/form", authController.RegisterForm)
    router.POST("/register", authController.HandleRegistration)

    log.Fatal(http.ListenAndServe(":8083", router))
}

