package routes

import (
    "log"
    "net/http"

    controller "github.com/lazarospsa/gophat/controller"
    "github.com/gorilla/mux"
)

func SetupAndRun() {
    mux := mux.NewRouter()
    
    mux.HandleFunc("/", controller.Index)
    mux.HandleFunc("/todo/{id}", controller.MarkTodo).Methods("PUT")
    mux.HandleFunc("/todo/{id}", controller.DeleteTodo).Methods("DELETE")
    mux.HandleFunc("/create", controller.CreateTodo).Methods("POST")

    log.Fatal(http.ListenAndServe(":5000", mux))
}












