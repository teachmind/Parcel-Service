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