package parcel

import (
	"context"
	"parcel-service/internal/app/model"
	svc "parcel-service/internal/app/service"

	"github.com/rs/zerolog/log"
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

// get Parcel List
func (s *service) GetParcels(ctx context.Context, status int, limit int, offset int) ([]model.Parcel, error) {
	parcel, err := s.repo.GetParcelsList(ctx, status, limit, offset)

	if err != nil {
		return nil, err
	}
	return parcel, nil
}

func (s *service) CreateParcel(ctx context.Context, parcel model.Parcel) error {
	const (
		CARRIER_FEE = 180.00
		COMPANY_FEE = 20.00
	)

	parcel.CarrierFee = CARRIER_FEE
	parcel.CompanyFee = COMPANY_FEE
	parcel.Price = CARRIER_FEE + COMPANY_FEE

	return s.repo.InsertParcel(ctx, parcel)
}

func (s *service) GetParcelByID(ctx context.Context, parcelID int) (model.Parcel, error) {
	parcel, err := s.repo.FetchParcelByID(ctx, parcelID)

	if err != nil {
		log.Error().Err(err).Msgf("[GetParcelByID] failed to get parcel Error: %v", err)
		return model.Parcel{}, err
	}

	return parcel, nil
}
