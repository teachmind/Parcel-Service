package parcel

import (
	"context"
	"parcel-service/internal/app/model"
	svc "parcel-service/internal/app/service"
)

type service struct {
	repo svc.ParcelRepository
}

func NewService(repo svc.ParcelRepository) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateParcel(ctx context.Context, parcel model.Parcel) (model.Parcel, error) {
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
	return s.repo.FetchParcelByID(ctx, parcelID)
}

func (s *service) EditParcel(ctx context.Context, parcel model.Parcel) error {
	return s.repo.UpdateParcel(ctx, parcel)
}
