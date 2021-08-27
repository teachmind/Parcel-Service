package carrier

import (
	"context"
	"parcel-service/internal/app/model"
	svc "parcel-service/internal/app/service"
	"time"
)
type service struct {
	repo svc.CarrierAcceptRepository
}

func NewService(repo svc.CarrierAcceptRepository) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) AssignCarrierToParcel(ctx context.Context, parcel model.CarrierRequest) error {
	status := model.Statuses{2, 3, 2}
	return s.repo.UpdateCarrierRequest(ctx, parcel, status.AcceptStatus, status.RejectStatus, status.ParcelStatus, time.Now())
}