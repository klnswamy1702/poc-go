// backend/controllers/controller.go
package controllers

import (
    "github.com/gin-gonic/gin"
    "github.com/klnswamy1702/poc-go-app/backend/models"
    "github.com/klnswamy1702/poc-go-app/backend/services"
    "net/http"
)

type controller struct {
    service *services.POCService
}

func Newcontroller(service *services.POCService) *controller {
    return &controller{service: service}
}

func (c *controller) CreatePOC(ctx *gin.Context) {
    var poc models.POC
    if err := ctx.BindJSON(&poc); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    err := c.service.CreatePOC(&poc)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, poc)
}

func (c *controller) GetAllPOCs(ctx *gin.Context) {
    pocs, err := c.service.GetAllPOCs()
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, pocs)
}

func (c *controller) GetPOCByID(ctx *gin.Context) {
    id := ctx.Param("id")
    poc, err := c.service.GetPOCByID(id)
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": "POC not found"})
        return
    }
    ctx.JSON(http.StatusOK, poc)
}

func (c *controller) UpdatePOC(ctx *gin.Context) {
    id := ctx.Param("id")
    var poc models.POC
    if err := ctx.BindJSON(&poc); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    err := c.service.UpdatePOC(id, &poc)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, poc)
}

func (c *controller) DeletePOC(ctx *gin.Context) {
    id := ctx.Param("id")
    err := c.service.DeletePOC(id)
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": "POC not found"})
        return
    }
    ctx.JSON(http.StatusOK, gin.H{"message": "POC deleted"})
}
