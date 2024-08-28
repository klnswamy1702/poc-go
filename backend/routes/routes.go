package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/klnswamy1702/poc-go/backend/controllers"
)

func SetupRouter(controller *controllers.controller) *gin.Engine {
    router := gin.Default()

    router.GET("/todos", controller.Gettodos)
    router.POST("/todos", controller.CreatePOC)
    router.PUT("/todos/:id", controller.UpdatePOC)
    router.DELETE("/todos/:id", controller.DeletePOC)

    return router
}
