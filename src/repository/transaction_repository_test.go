package repository

import (
	"pismo-transactions/src/models"
	"testing"
	"time"
)

func TestInsertTransactionRepository(t *testing.T) {
	type args struct {
		t models.Transaction
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Sucess",
			args: args{
				models.Transaction{
					TransactionID:   1,
					AccountID:       3,
					OperationTypeID: 4,
					Amount:          -200,
					EventDate:       time.Now().Format("01-06-2006 00:00:00"),
				},
			},
			wantErr: false,
		},
		{
			name: "Fail",
			args: args{
				models.Transaction{
					AccountID:       0,
					TransactionID:   0,
					OperationTypeID: 0,
					Amount:          0.0,
					EventDate:       time.Now().Format("01-06-2006 00:00:00"),
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := InsertTransactionRepository(tt.args.t); (err != nil) != tt.wantErr {
				t.Errorf("InsertTransactionRepository() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
