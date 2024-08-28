package main

import (
    "github.com/klnswamy1702/poc-go/backend/config"
    "github.com/klnswamy1702/poc-go/backend/controllers"
    "github.com/klnswamy1702/poc-go/backend/repositories"
    "github.com/klnswamy1702/poc-go/backend/routes"
    "github.com/klnswamy1702/poc-go/backend/services"
)

func main() {
    config.ConnectDB()

    Repo := repositories.Newrepository()
    Service := services.Newservice(repo)
    Controller := controllers.Newcontroller(service)

    router := routes.SetupRouter(controller)
    router.Run(":8080")
}
