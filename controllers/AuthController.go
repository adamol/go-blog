package controllers

import (
    "net/http"
    "github.com/julienschmidt/httprouter"
    "blog/models"
    "github.com/satori/go.uuid"
)

var sessions = map[string]string{}

// func alreadyLoggedIn(req *http.Request) bool {
//     c, err := req.Cookie("session")
//     if err != nil {
//         return false
//     }
//     un := sessions[c.Value]
//     _, ok := models.User{}.Find(un)
//     return ok
// }

type AuthController struct {}

func (ac AuthController) LoginForm (res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
    view(res, "auth.login", "")
}

func (ac AuthController) HandleLogin (res http.ResponseWriter, req *http.Request, params httprouter.Params) {
    req.ParseForm()
    username := req.Form.Get("username")
    password := req.Form.Get("password")

    user := models.User{}.Find(username)

    err := bcrypt.CompareHashAndPassword(user.Password, []byte(password))

    if err != nil {
        http.Error(res, "User and/or password do not match", http.StatusForbidden)
        return
    }

    sID := uuid.NewV4()
    c := &http.Cookie{
        Name: "session",
        Value: sID.String(),
    }
    http.SetCookie(res, c)
    sessions[c.Value] = username

    http.Redirect(res, req, "/", http.StatusSeeOther)
}

func (ac AuthController) RegisterForm (res http.ResponseWriter, req *http.Request, params httprouter.Params) {
    view(res, "auth.register", "")
}

func (ac AuthController) HandleRegistration (res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
    //if alreadyLoggedIn(req) {
    //    http.Redirect(w, req, "/", http.StatusSeeOther)
    //    return
    //}

    req.ParseForm()
    username := req.Form.Get("username")
    password := req.Form.Get("password")

    hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
    if err != nil {
        http.Error(res, "Internal server error", http.StatusInternalServerError)
        return
    }

    models.User{}.Create(username, hashedPass)

    // @TODO: implement error handling
    // if alreadyLoggedIn(req) {
    //     http.Redirect(w, req, "/", http.StatusSeeOther)
    //     return
    // }

    sID := uuid.NewV4()
    c := &http.Cookie {
        Name: "session",
        Value: sID.String(),
    }
    http.SetCookie(res, c)
    sessions[c.Value] = username

    http.Redirect(res, req, "/", 301)
}

