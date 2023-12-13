package main

import (
    "github.com/lazarospsa/gophat/model"
    "github.com/lazarospsa/gophat/routes"
)

func main() {

    model.Setup()
    routes.SetupAndRun()

}
