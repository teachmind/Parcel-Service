package service

import (
	"context"
	"parcel-service/internal/app/model"
)

// ParcelRepository to Insert New Parcel
type ParcelRepository interface {
	InsertParcel(ctx context.Context, parcel model.Parcel) error
	FetchParcelByID(ctx context.Context, parcelID int) (model.Parcel, error)
	UpdateParcel(ctx context.Context, parcel model.Parcel) error
}

// ParcelService to Create new parcel
type ParcelService interface {
	CreateParcel(ctx context.Context, parcel model.Parcel) error
	GetParcelByID(ctx context.Context, parcelID int) (model.Parcel, error)
	EditParcel(ctx context.Context, parcel model.Parcel) error
}

// CarrierRepository to update the carrier request table
type CarrierRepository interface {
	InsertCarrierRequest(ctx context.Context, carrierReq model.CarrierRequest) error
}

// CarrierService to add new carrier request
type CarrierService interface {
	NewCarrierRequest(ctx context.Context, carrierReq model.CarrierRequest) error
}
