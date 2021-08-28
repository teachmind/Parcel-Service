package service

import (
	"context"
	"parcel-service/internal/app/model"
)

// ParcelRepository to Insert New Parcel and get parcel list
type ParcelRepository interface {
	InsertParcel(ctx context.Context, parcel model.Parcel) error
	GetParcelsList(ctx context.Context, status int, limit int, offset int) ([]model.Parcel, error)
}

// ParcelService to Create new parcel & get parcel list
type ParcelService interface {
	CreateParcel(ctx context.Context, parcel model.Parcel) error
	GetParcels(ctx context.Context, status int, limit int, offset int) ([]model.Parcel, error)
}
