package router

import (
	"net/http"
	"pismo-transactions/src/errs"
	"pismo-transactions/src/models"
	"pismo-transactions/src/services"

	"github.com/labstack/echo/v4"
)

// @Summary      Create new Transaction
// @Description  POST models.Transaction
// @Tags         transaction
// @Accept       json
// @Produce      json
// @Param        Transaction   body     models.Transaction  true  "Transaction"
// @Success      201
// @Failure      400  {object}  errs.Handling
// @Router       /transaction [post]
func CreateTransaction(c echo.Context) error {
	var transaction models.Transaction
	if err := c.Bind(&transaction); err != nil {
		eh := errs.Handling{Error: err.Error(), Message: errs.ERR_BIND_OBJECT}
		return c.JSON(http.StatusBadRequest, eh)
	}

	err := services.CreateTransactionService(transaction)
	if err != nil {
		eh := errs.Handling{Error: err.Error(), Message: errs.ERR_SERVICE}
		return c.JSON(http.StatusBadRequest, eh)
	}

	return c.NoContent(http.StatusCreated)
}
