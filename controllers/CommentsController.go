package controllers

import (
    "net/http"

    "github.com/julienschmidt/httprouter"

    "blog/models"
)

type CommentsController struct {}

//func (pc PostsController) Index (res http.ResponseWriter, req *http.Request, params httprouter.Params) {
//    comments := models.Comment{}.List()
//
//    view(res, "comments.index", comments)

func (cc CommentsController) Store (res http.ResponseWriter, req *http.Request, params httprouter.Params) {
    req.ParseForm()
    body  := req.Form.Get("body")
    slug  := params.ByName("slug")

    post := models.Post{}.Find(slug)
    post.Comments = append(post.Comments, models.Comment{body})
    post.Update()

    //view(res, "posts.show", post)
    //models.Post{}.AddComment(slug, body)

    http.Redirect(res, req, "/posts/" + slug, 301)
}

