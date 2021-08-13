package service

import (
	"context"
	"parcel-service/internal/app/model"

)

//go:generate mockgen -source service.go -destination ./mocks/mock_service.go -package mocks
// CarrierParcelAcceptRepository to update the carrier request table
type CarrierParcelAcceptRepository interface {
	UpdateCarrierRequest(ctx context.Context, parcel model.CarrierRequest, status model.ParcelStatus) error
}

// UserService to fetch user by PhoneNumber and Password
type CarrierParcelAcceptService interface {
	AssignCarrierToParcel(ctx context.Context, parcel model.CarrierRequest, status model.ParcelStatus) error
}