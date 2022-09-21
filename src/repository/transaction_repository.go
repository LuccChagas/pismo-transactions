package repository

import (
	"fmt"
	"pismo-transactions/src/database"
	"pismo-transactions/src/errs"
	"pismo-transactions/src/models"
)

func InsertTransactionRepository(t models.Transaction) error {
	Conn, err := database.Conn()
	if err != nil {
		return fmt.Errorf("%s %w", errs.ERR_DB_CONN, err)
	}

	query := `INSERT INTO public.transactions (account_id, operation_type_id, amount, event_date)
	 VALUES($1, $2, $3, $4);
	`

	_, err = Conn.Exec(query, t.AccountID, t.OperationTypeID, t.Amount, t.EventDate)
	if err != nil {
		return fmt.Errorf("%s %w", errs.ERR_EXEC_QUERY, err)
	}

	defer Conn.Close()
	return nil
}
