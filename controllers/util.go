package controllers

import (
    "html/template"
    "net/http"
    "strings"
)

func view (res http.ResponseWriter, path string, params interface{}) {
    s := strings.Split(path, ".")
    folder, file := s[0], s[1]

    folder = "resources/views/" + folder + "/*"
    file   = file + ".gohtml"

    var tpl *template.Template
    tpl = template.Must(template.ParseGlob(folder))

    tpl.ExecuteTemplate(res, file, params)
}

