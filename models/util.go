package models

import (
    "gopkg.in/mgo.v2"
)

func getDbHandler() *mgo.Session {
    mongo, err := mgo.Dial("localhost")
    if err != nil {
        panic(err)
    }
    return mongo
}

func getPostsRepo() *mgo.Collection {
    mongo, err := mgo.Dial("localhost")
    if err != nil {
        panic(err)
    }

    return mongo.DB("blog").C("posts")
}

func getCommentsRepo() *mgo.Collection {
    mongo, err := mgo.Dial("localhost")
    if err != nil {
        panic(err)
    }

    return mongo.DB("blog").C("comments")
}

