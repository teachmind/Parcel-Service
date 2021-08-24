package model

import (
	"fmt"
	"time"
)

type Parcel struct {
	ID                 int       `json:"id"`
	UserID             int       `json:"user_id" db:"user_id"`
	CarrierID          int       `json:"carrier_id" db:"carrier_id"`
	Status             int       `json:"status"`
	SourceAddress      string    `json:"source_address" db:"source_address"`
	DestinationAddress string    `json:"destination_address" db:"destination_address"`
	SourceTime         time.Time `json:"source_time" db:"source_time"`
	ParcelType         string    `json:"type" db:"type"`
	Price              float32   `json:"price" db:"price"`
	CarrierFee         float32   `json:"carrier_fee" db:"carrier_fee"`
	CompanyFee         float32   `json:"company_fee" db:"company_fee"`
}

type CarrierRequest struct {
	ID        int `json:"id"`
	ParcelID  int `json:"parcel_id" db:"parcel_id"`
	CarrierID int `json:"carrier_id" db:"carrier_id"`
	Status    int `json:"status" db:"status"`
}

func (p *Parcel) ValidateParcelInput() error {
	if p.SourceAddress == "" {
		return fmt.Errorf("source Address is required :%w", ErrEmpty)
	}

	if p.DestinationAddress == "" {
		return fmt.Errorf("destination Address is required :%w", ErrEmpty)
	}

	if p.ParcelType == "" {
		return fmt.Errorf("Parcel type is required :%w", ErrEmpty)
	}

	if p.UserID == 0 {
		return fmt.Errorf("user ID is required :%w", ErrEmpty)
	}

	if !p.SourceTime.IsZero() && p.SourceTime.Before(time.Now()) {
		return fmt.Errorf("source time must be future date:%w", ErrEmpty)
	}

	return nil
}

// Validates carrier request input credentials
func (cr *CarrierRequest) ValidateCarrierId() error {
	if cr.CarrierID == 0 {
		return fmt.Errorf("Carrier ID is required :%w", ErrEmpty)
	}
	return nil
}
