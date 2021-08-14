package parcel_carrier

import (
	"context"
	"reflect"
	"testing"
)

func TestNewRepository(t *testing.T) {
	type args struct {
		db *sqlx.DB
	}
	tests := []struct {
		name string
		args args
		want *repository
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRepository(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repository_UpdateCarrierRequest(t *testing.T) {
	type fields struct {
		db *sqlx.DB
	}
	type args struct {
		ctx    context.Context
		parcel CarrierRequest
		status ParcelStatus
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository{
				db: tt.fields.db,
			}
			if err := r.UpdateCarrierRequest(tt.args.ctx, tt.args.parcel, tt.args.status); (err != nil) != tt.wantErr {
				t.Errorf("UpdateCarrierRequest() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
