package repository

import (
	"pismo-transactions/src/models"
	"reflect"
	"testing"
)

func TestInsertAccount(t *testing.T) {
	type args struct {
		account models.Account
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Success",
			args: args{
				models.Account{
					AccountID:      1,
					DocumentNumber: "123445566778",
				},
			},
			wantErr: false,
		},
		{
			name: "Fail",
			args: args{
				models.Account{
					AccountID:      0,
					DocumentNumber: "",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := InsertAccount(tt.args.account); (err != nil) != tt.wantErr {
				t.Errorf("InsertAccount() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetAccountByID(t *testing.T) {
	type args struct {
		ID string
	}
	tests := []struct {
		name    string
		args    args
		want    *models.Account
		wantErr bool
	}{
		{
			name: "Success",
			args: args{
				ID: "1",
			},
			want: &models.Account{
				AccountID:      1,
				DocumentNumber: "12345678900",
			},
			wantErr: false,
		},
		{
			name: "Fail",
			args: args{
				ID: "",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAccountByID(tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAccountByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAccountByID() = %v, want %v", got, tt.want)
			}
		})
	}
}
