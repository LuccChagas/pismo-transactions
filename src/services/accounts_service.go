package services

import (
	"pismo-transactions/src/models"
	"pismo-transactions/src/repository"
	"regexp"
	"strings"
)

func CreateAccountService(account models.Account) error {

	regex := regexp.MustCompile("[^0-9]")
	account.DocumentNumber = regex.ReplaceAllString(account.DocumentNumber, "")
	account.DocumentNumber = strings.TrimSpace(account.DocumentNumber)

	err := repository.InsertAccount(account)
	if err != nil {
		return err
	}

	return nil
}

func GetAccountByIDService(id string) (*models.Account, error) {

	account, err := repository.GetAccountByID(id)
	if err != nil {
		return nil, err
	}

	return account, nil
}
