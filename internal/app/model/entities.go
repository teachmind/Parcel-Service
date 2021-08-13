package model

import "fmt"

type CarrierRequest struct {
	ID          int     `json:"id"`
	ParcelID 	int 	`json:"parcel_id" db:"parcel_id"`
	CarrierID   int 	`json:"carrier_id" db:"carrier_id"`
	status      int     `json:"status" db:"status"`
}

// Validates carrier request input credentials
func (cr *CarrierRequest) ValidateCarrierId() error {
	if cr.CarrierID == 0 {
		return fmt.Errorf("Carrier ID is required :%w", ErrEmpty)
	}
	return nil
}