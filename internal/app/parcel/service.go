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

// get Parcel List
func (s *service) GetParcels(ctx context.Context, limit int, offset int) ([]model.Parcel, error) {
	parcel, err := s.repo.SelectParcelsList(ctx, limit, offset)

	if err != nil {
		return nil, err
	}

	return parcel, nil
}
