package models

//import (
//    "log"
//)

type Comment struct {
    Body string
}

// func (comment Comment) Create(postSlug string, commentBody string) {
//     dbh := getDbHandler()
//     posts := dbh.DB("blog").C("posts")
//     post := posts.Find(postSlug)
//
//     post.comments = append(post.comments(&Comment{body}))
//
//     posts.Update(postSlug, post)
//
//     if err != nil {
//         log.Fatal(err)
//     }
// }

