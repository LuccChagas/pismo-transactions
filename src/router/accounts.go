package router

import (
	"net/http"
	"pismo-transactions/src/errs"
	"pismo-transactions/src/models"
	"pismo-transactions/src/services"

	"github.com/labstack/echo/v4"
)

// @Summary      Create new Account
// @Description  POST models.Account
// @Tags         Accounts
// @Accept       json
// @Produce      json
// @Param        Account   body     models.Account  true  "Account"
// @Success      201
// @Failure      400  {object}  errs.Handling
// @Router       /account [post]
func CreateAccount(c echo.Context) error {
	var account models.Account
	if err := c.Bind(&account); err != nil {
		eh := errs.Handling{Error: err.Error(), Message: errs.ERR_BIND_OBJECT}
		return c.JSON(http.StatusBadRequest, eh)
	}

	err := services.CreateAccountService(account)
	if err != nil {
		eh := errs.Handling{Error: err.Error(), Message: errs.ERR_SERVICE}
		return c.JSON(http.StatusBadRequest, eh)
	}

	return c.NoContent(http.StatusCreated)
}

// @Summary      Get Account By ID
// @Description  GET models.Account
// @Tags         Accounts
// @Accept       json
// @Produce      json
// @Param        id   query      string  true  "Account ID"
// @Success      200  {object}	[]models.Account
// @Failure      400  {object}  errs.Handling
// @Router       /account:id [get]
func GetAccountByID(c echo.Context) error {
	id := c.QueryParam("accountId")
	if len(id) <= 0 {
		return c.JSON(http.StatusBadRequest, "accountId is required")
	}

	accounts, err := services.GetAccountByIDService(id)
	if err != nil {
		eh := errs.Handling{Error: err.Error(), Message: errs.ERR_SERVICE}
		return c.JSON(http.StatusBadRequest, eh)
	}

	return c.JSON(http.StatusOK, accounts)
}
