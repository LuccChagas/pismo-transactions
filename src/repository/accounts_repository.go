package repository

import (
	"fmt"
	"pismo-transactions/src/database"
	"pismo-transactions/src/errs"
	"pismo-transactions/src/models"
)

func InsertAccount(account models.Account) error {
	Conn, err := database.Conn()
	if err != nil {
		return fmt.Errorf("%s %w", errs.ERR_DB_CONN, err)
	}

	if len(account.DocumentNumber) < 1 {
		return fmt.Errorf("DocumentNumber don't exists")
	}

	query := `INSERT INTO public.accounts (document_number) VALUES($1);
	`

	_, err = Conn.Exec(query, account.DocumentNumber)
	if err != nil {
		return fmt.Errorf("%s %w", errs.ERR_EXEC_QUERY, err)
	}

	defer Conn.Close()
	return nil
}

func GetAccountByID(ID string) (*models.Account, error) {
	Conn, err := database.Conn()
	if err != nil {
		return nil, fmt.Errorf("%s %w", errs.ERR_DB_CONN, err)
	}

	if len(ID) < 1 {
		return nil, fmt.Errorf("AccountID don't exists")
	}

	query := `SELECT account_id, document_number FROM public.accounts WHERE account_id = $1;
	`
	rows, err := Conn.Query(query, ID)
	if err != nil {
		return nil, fmt.Errorf("%s %w", errs.ERR_EXEC_QUERY, err)
	}

	account := new(models.Account)
	for rows.Next() {
		if err := rows.Scan(
			&account.AccountID,
			&account.DocumentNumber,
		); err != nil {
			return nil, err
		}
	}

	defer Conn.Close()
	return account, nil
}
