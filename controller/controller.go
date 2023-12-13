package controller

import (
    "net/http"
    "fmt"
    "strconv"

    "html/template"
    "github.com/lazarospsa/gophat/model"

    "github.com/gorilla/mux"
)

func sendTodos(w http.ResponseWriter) {

    todos, err := model.GetAllTodos()
    if err != nil {
        fmt.Println("Could not get all todos from db", err)
        return
    }

    tmpl := template.Must(template.ParseFiles("templates/index.html"))

    err = tmpl.ExecuteTemplate(w, "Todos", todos)
    if err != nil {
        fmt.Println("Could not execute template", err)
    }
}


func Index(w http.ResponseWriter, r *http.Request) {
    
    todos, err := model.GetAllTodos()
    if err != nil {
        fmt.Println("Could not get all todos from db", err)
    }

    tmpl := template.Must(template.ParseFiles("templates/index.html"))

    err = tmpl.Execute(w, todos)
    if err != nil {
        fmt.Println("Could not execute template", err)
    }

}


func MarkTodo(w http.ResponseWriter, r *http.Request) {

    id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
    if err != nil {
        fmt.Println("Could not parse id", err)
    }

    err = model.MarkDone(id)
    if err != nil {
        fmt.Println("Could not update todo", err)
    }



    sendTodos(w)



}

func CreateTodo(w http.ResponseWriter, r *http.Request) {

    err := r.ParseForm()
    if err != nil {
        fmt.Println("Error parsing form", err)
    }

    err = model.CreateTodo(r.FormValue("todo"))
    if err != nil {
        fmt.Println("Could not create todo", err)
    }

    sendTodos(w)

}


func DeleteTodo(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
    if err != nil {
        fmt.Println("Could not parse id", err)
    }

    err = model.Delete(id)
    if err != nil {
        fmt.Println("Could not delete", err)
    }

    sendTodos(w)

}