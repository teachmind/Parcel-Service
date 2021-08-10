package model


type CarrierRequest struct {
	ID          int     `json:"id"`
	ParcelID 	int 	`json:"parcel_id" db:"parcel_id"`
	CarrierID   int 	`json:"carrier_id" db:"carrier_id"`
	status      int     `json:"status" db:"status"`
}
