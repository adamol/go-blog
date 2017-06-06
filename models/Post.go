package models

import (
    "log"
    "gopkg.in/mgo.v2/bson"
)

type Post struct {
    Title string
    Body string
    Slug string
    Comments []Comment
}

func (post Post) List() []Post {
    dbh := getDbHandler()
    posts := dbh.DB("blog").C("posts")

    var results []Post

    err := posts.Find(bson.M{}).All(&results)
    if err != nil {
        log.Fatal(err)
    }

    return results
}

func (post Post) Create(title string, body string, slug string) {
    dbh := getDbHandler()
    posts := dbh.DB("blog").C("posts")

    comments := []Comment{}
    err := posts.Insert(&Post{title, body, slug, comments})

    if err != nil {
        log.Fatal(err)
    }
}

func (post Post) Find(slug string) Post {
    dbh := getDbHandler()
    posts := dbh.DB("blog").C("posts")

    var result Post

    err := posts.Find(bson.M{"slug": slug}).One(&result)
    if err != nil {
        log.Fatal(err)
    }

    return result
}

func (p Post) Update() {
    dbh := getDbHandler()
    posts := dbh.DB("blog").C("posts")

    query := bson.M{"slug": p.Slug}
    update := bson.M{"$set": &Post{p.Title, p.Body, p.Slug, p.Comments}}

    err := posts.Update(query, update)

    if err != nil {
        log.Fatal(err)
    }
}
