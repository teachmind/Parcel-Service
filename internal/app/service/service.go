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

// CarrierParcelRequestRepository to update the carrier request table
type CarrierRequestRepository interface {
	InsertCarrierRequest(ctx context.Context, carrierReq model.CarrierRequest) error
}

// CarrierRequestService to add new carrier request
type CarrierRequestService interface {
	AddCarrierReqest(ctx context.Context, carrierReq model.CarrierRequest) error
}
