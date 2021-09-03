//go:generate mockgen -source=internal/app/service/service.go -destination=internal/app/service/mocks/mock_service.go

package service

import (
	"context"
	"parcel-service/internal/app/model"
)

// ParcelRepository to Insert New Parcel and get parcel list
type ParcelRepository interface {
	InsertParcel(ctx context.Context, parcel model.Parcel) error
	FetchParcelByID(ctx context.Context, parcelID int) (model.Parcel, error)
	GetParcelsList(ctx context.Context, status int, limit int, offset int) ([]model.Parcel, error)
	UpdateParcel(ctx context.Context, parcel model.Parcel) error
}

// ParcelService to Create new parcel & get parcel list
type ParcelService interface {
	CreateParcel(ctx context.Context, parcel model.Parcel) error
	GetParcelByID(ctx context.Context, parcelID int) (model.Parcel, error)
	GetParcels(ctx context.Context, status int, limit int, offset int) ([]model.Parcel, error)
	EditParcel(ctx context.Context, parcel model.Parcel) error
}

type CarrierRepository interface {
	InsertCarrierRequest(ctx context.Context, carrierReq model.CarrierRequest) error
}

type CarrierService interface {
	NewCarrierRequest(ctx context.Context, carrierReq model.CarrierRequest) error
}
