package service

import (
	"context"
	"parcel-service/internal/app/model"
)

// ParcelRepository to Insert New Parcel
type ParcelRepository interface {
	InsertParcel(ctx context.Context, parcel model.Parcel) error
}

// ParcelService to Create new parcel
type ParcelService interface {
	CreateParcel(ctx context.Context, parcel model.Parcel) error
}
