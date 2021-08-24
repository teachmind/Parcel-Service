package service

import (
	"context"
	"parcel-service/internal/app/model"
	"time"
)

//go:generate mockgen -source service.go -destination ./mocks/mock_service.go -package mocks
type CarrierAcceptRepository interface {
	UpdateCarrierRequest(ctx context.Context, parcel model.CarrierRequest, acceptStatus int, rejectStatus int, parcelStatus int, sourceTime time.Time) error
}

type CarrierAcceptService interface {
	AssignCarrierToParcel(ctx context.Context, parcel model.CarrierRequest) error
}

// ParcelRepository to Insert New Parcel
type ParcelRepository interface {
	InsertParcel(ctx context.Context, parcel model.Parcel) error
}

// ParcelService to Create new parcel
type ParcelService interface {
	CreateParcel(ctx context.Context, parcel model.Parcel) error
}

