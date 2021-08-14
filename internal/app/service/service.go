package service

import (
	"context"
	"parcel-service/internal/app/model"
)

// ParcelRepository to fetch parcel by PhoneNumber
type ParcelRepository interface {
	InsertParcel(ctx context.Context, parcel model.Parcel) error
}

// ParcelService to fetch parcel by PhoneNumber and Password
type ParcelService interface {
	CreateParcel(ctx context.Context, parcel model.Parcel) error
}
