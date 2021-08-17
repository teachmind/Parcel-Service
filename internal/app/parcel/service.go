package parcel

import (
	"context"
	"parcel-service/internal/app/model"
	svc "parcel-service/internal/app/service"
)

type service struct {
	repo svc.ParcelRepository
}

// NewService is to generate for new repo
func NewService(repo svc.ParcelRepository) *service {
	return &service{
		repo: repo,
	}
}

// CreaetParcel creates new parcel
func (s *service) CreateParcel(ctx context.Context, parcel model.Parcel) error {
	return s.repo.InsertParcel(ctx, parcel)
}
