package model

import "fmt"

type Parcel struct {
	ID                 int     `json:"id"`
	UserID             int     `json:"user_id" db:"user_id"`
	CarrierID          int     `json:"carrier_id" db:"carrier_id"`
	Status             int     `json:"status"`
	SourceAddress      string  `json:"source_address" db:"source_address"`
	DestinationAddress string  `json:"destination_address" db:"destination_address"`
	SourceTime         string  `json:"source_time" db:"source_time"`
	ParcelType         string  `json:"type" db:"type"`
	Price              float32 `json:"price" db:"price"`
	CarrierFee         float32 `json:"carrier_fee" db:"carrier_fee"`
	CompanyFee         float32 `json:"company_fee" db:"company_fee"`
}

func (p *Parcel) ValidateParcelInput() error {
	if p.SourceAddress == "" {
		return fmt.Errorf("Source Address is required :%w", ErrEmpty)
	}

	if p.DestinationAddress == "" {
		return fmt.Errorf("Destination Address is required :%w", ErrEmpty)
	}

	if p.ParcelType == "" {
		return fmt.Errorf("Parcel Type is required :%w", ErrEmpty)
	}

	if p.UserID == 0 {
		return fmt.Errorf("User ID is required :%w", ErrEmpty)
	}

	return nil
}
