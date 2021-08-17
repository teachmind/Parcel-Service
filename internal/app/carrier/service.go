package carrier

import (
	"context"
	"parcel-service/internal/app/model"
	svc "parcel-service/internal/app/service"
)

type service struct {
	repo svc.CarrierRequestRepository
}

// NewService is to generate for new repo
func NewService(repo svc.CarrierRequestRepository) *service {
	return &service{
		repo: repo,
	}
}

// Add a new carrier request to
func (s *service) NewCarrierRequest(ctx context.Context, carrierReq model.CarrierRequest) error {
	return s.repo.InsertCarrierRequest(ctx, carrierReq)
}
