package parcel_carrier

import (
	"context"
	"parcel-service/internal/app/model"
	svc "parcel-service/internal/app/service"
	"time"
)
type service struct {
	repo svc.CarrierParcelAcceptRepository
}


// NewService is to generate for new repo
func NewService(repo svc.CarrierParcelAcceptRepository) *service {
	return &service{
		repo: repo,
	}
}

// Assign a parcel with a carrier is to accept a carrier and other request will be rejected
func (s *service) AssignCarrierToParcel(ctx context.Context, parcel model.CarrierRequest) error {
	return s.repo.UpdateCarrierRequest(ctx, parcel, 2, 3, 2, time.Now())
}