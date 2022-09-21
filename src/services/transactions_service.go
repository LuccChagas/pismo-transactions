package services

import (
	"pismo-transactions/src/models"
	"pismo-transactions/src/repository"
	"time"
)

func CreateTransactionService(t models.Transaction) error {

	t.EventDate = time.Now().Format("01-06-2006 00:00:00")
	err := repository.InsertTransactionRepository(t)
	if err != nil {
		return err
	}

	return nil
}
