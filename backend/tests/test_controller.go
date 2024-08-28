package tests

import (
    "testing"

    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
    "net/http"
    "net/http/httptest"

    "github.com/klnswamy1702/poc-go/backend/controllers"
    "github.com/klnswamy1702/poc-go/backend/models"
    "github.com/klnswamy1702/poc-go/backend/services"
)

func TestGettodos(t *testing.T) {
    gin.SetMode(gin.TestMode)

    w := httptest.NewRecorder()
    c, _ := gin.CreateTestContext(w)

    mockService := new(services.Mocktodoservice)
    mockService.On("GetAlltodos").Return([]models.POC{}, nil)

    controller := controllers.Newcontroller(mockService)
    controller.Gettodos(c)

    assert.Equal(t, http.StatusOK, w.Code)
}

func TestCreatePOC(t *testing.T) {
    gin.SetMode(gin.TestMode)

    w := httptest.NewRecorder()
    c, _ := gin.CreateTestContext(w)

    poc := models.POC{
        Title:       "Test POC",
        Description: "Test Description",
        Completed:   false,
    }

    c.Request, _ = http.NewRequest("POST", "/todos", nil)
    c.Request.Header.Set("Content-Type", "application/json")

    mockService := new(services.Mocktodoservice)
    mockService.On("CreatePOC", poc).Return(nil)

    controller := controllers.Newcontroller(mockService)
    controller.CreatePOC(c)

    assert.Equal(t, http.StatusCreated, w.Code)
}

func TestUpdatePOC(t *testing.T) {
    gin.SetMode(gin.TestMode)

    w := httptest.NewRecorder()
    c, _ := gin.CreateTestContext(w)

    poc := models.POC{
        Title:       "Updated POC",
        Description: "Updated Description",
        Completed:   true,
    }

    c.Params = []gin.Param{
        {Key: "id", Value: "some-id"},
    }

    c.Request, _ = http.NewRequest("PUT", "/todos/some-id", nil)
    c.Request.Header.Set("Content-Type", "application/json")

    mockService := new(services.Mocktodoservice)
    mockService.On("UpdatePOC", "some-id", poc).Return(nil)

    controller := controllers.Newcontroller(mockService)
    controller.UpdatePOC(c)

    assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeletePOC(t *testing.T) {
    gin.SetMode(gin.TestMode)

    w := httptest.NewRecorder()
    c, _ := gin.CreateTestContext(w)

    c.Params = []gin.Param{
        {Key: "id", Value: "some-id"},
    }

    c.Request, _ = http.NewRequest("DELETE", "/todos/some-id", nil)

    mockService := new(services.Mocktodoservice)
    mockService.On("DeletePOC", "some-id").Return(nil)

    controller := controllers.Newcontroller(mockService)
    controller.DeletePOC(c)

    assert.Equal(t, http.StatusOK, w.Code)
}
