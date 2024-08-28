// backend/services/service.go
package services

import (
    "github.com/klnswamy1702/poc-go-app/backend/models"
    "github.com/klnswamy1702/poc-go-app/backend/repositories"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type service struct {
    repo *repositories.repositories
}

func Newservice(repo *repositories.repositories) *service {
    return &service{repo: repo}
}

func (s *service) CreatePOC(poc *models.POC) error {
    return s.repo.CreatePOC(poc)
}

func (s *service) GetAllPOCs() ([]models.POC, error) {
    return s.repo.GetAllPOCs()
}

func (s *service) GetPOCByID(id string) (*models.POC, error) {
    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return nil, err
    }
    return s.repo.GetPOCByID(objID)
}

func (s *service) UpdatePOC(id string, poc *models.POC) error {
    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return err
    }
    return s.repo.UpdatePOC(objID, poc)
}

func (s *service) DeletePOC(id string) error {
    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return err
    }
    return s.repo.DeletePOC(objID)
}
