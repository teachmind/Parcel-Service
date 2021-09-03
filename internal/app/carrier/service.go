package carrier

import (
	"context"
	"parcel-service/internal/app/model"
	svc "parcel-service/internal/app/service"
)

type service struct {
	repo svc.CarrierRepository
}

func NewService(repo svc.CarrierRepository) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) NewCarrierRequest(ctx context.Context, carrierReq model.CarrierRequest) error {
	return s.repo.InsertCarrierRequest(ctx, carrierReq)
}
