package controllers

import (
    "net/http"
    "strings"

    "github.com/julienschmidt/httprouter"

    "blog/models"
)

type PostsController struct {}

func (pc PostsController) Index (res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
    posts := models.Post{}.List()

    view(res, "posts.index", posts)
}

func (pc PostsController) Store (res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
    req.ParseForm()
    title := req.Form.Get("title")
    body  := req.Form.Get("body")
    slug  := strings.ToLower(strings.Replace(title, " ", "-", -1))

    models.Post{}.Create(title, body, slug)

    http.Redirect(res, req, "/", 301)
}

func (pc PostsController) Show (res http.ResponseWriter, req *http.Request, params httprouter.Params) {
    slug := params.ByName("slug")

    post := models.Post{}.Find(slug)

    view(res, "posts.show", post)
}

