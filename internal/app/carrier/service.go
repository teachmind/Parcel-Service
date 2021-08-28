package carrier

import (
	"context"
	"parcel-service/internal/app/model"
	svc "parcel-service/internal/app/service"
	"time"
)

type service struct {
	repo svc.CarrierRepository
}

func NewService(repo svc.CarrierRepository) *service {
	return &service{
		repo: repo,
	}
}

// Add a new carrier request to
func (s *service) NewCarrierRequest(ctx context.Context, carrierReq model.CarrierRequest) error {
	return s.repo.InsertCarrierRequest(ctx, carrierReq)
}

func (s *service) AssignCarrierToParcel(ctx context.Context, parcel model.CarrierRequest) error {
	status := model.Statuses{2, 3, 2}
	return s.repo.UpdateCarrierRequest(ctx, parcel, status.AcceptStatus, status.RejectStatus, status.ParcelStatus, time.Now())
}
