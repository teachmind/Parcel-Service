package service

import (
	"context"
	"parcel-service/internal/app/model"
)


// CarrierParcelAcceptRepository to update the carrier request table
type CarrierParcelAcceptRepository interface {
	AssignCarrierToParcel(ctx context.Context, parcel model.CarrierRequest) error
}

// UserService to fetch user by PhoneNumber and Password
type CarrierParcelAcceptService interface {
	AssignCarrierToParcel(ctx context.Context, parcel model.CarrierRequest) error
}