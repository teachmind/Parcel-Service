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

const acceptStatus, rejectStatus, parcelStatus int = 2, 3, 2

func NewService(repo svc.CarrierRepository) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) NewCarrierRequest(ctx context.Context, carrierReq model.CarrierRequest) error {
	return s.repo.InsertCarrierRequest(ctx, carrierReq)
}

func (s *service) AssignCarrierToParcel(ctx context.Context, parcel model.CarrierRequest) error {
	return s.repo.UpdateCarrierRequest(ctx, parcel, acceptStatus, rejectStatus, parcelStatus, time.Now())
}
