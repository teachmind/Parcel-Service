package model

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
	CreatedAt          string  `json:"created_at" db:"created_at"`
	UpdatedAt          string  `json:"updated_at" db:"updated_at"`
}
