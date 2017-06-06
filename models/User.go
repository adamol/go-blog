package models

import (
    "log"
    "gopkg.in/mgo.v2/bson"
)

type User struct {
    Username string
    Password []byte
}

func (user User) Create(username string, password []byte) {
    dbh := getDbHandler()
    users := dbh.DB("blog").C("users")

    err := users.Insert(&User{username, password})

    if err != nil {
        log.Fatal(err)
    }
}

func (user User) Find(username string) User {
    dbh := getDbHandler()
    users := dbh.DB("blog").C("users")

    var result User

    err := users.Find(bson.M{"username": username}).One(&result)
    if err != nil {
        log.Fatal(err)
    }

    return result
}

