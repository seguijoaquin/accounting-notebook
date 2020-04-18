package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/seguijoaquin/accounting-notebook/cmd/api/apierror"
	"github.com/seguijoaquin/accounting-notebook/internal/domain"
	"github.com/seguijoaquin/accounting-notebook/internal/errorx"
	"github.com/seguijoaquin/accounting-notebook/internal/processor"
)

//Transaction represents a transaction operation from an account
type Transaction interface {
	HandleNewTransaction(c *gin.Context)
	HandleFindTransactionByID(c *gin.Context)
	HandleFetchAllTransactions(c *gin.Context)
}

type transaction struct {
	p processor.Transaction
}

func (t *transaction) HandleNewTransaction(c *gin.Context) {

	var newTransaction domain.TransactionBody

	if err := c.ShouldBindJSON(&newTransaction); err != nil {
		e := apierror.NewBadRequest(fmt.Sprintf("invalid input: %s", err.Error()))
		c.JSON(
			e.Status(),
			e.JSON(),
		)
		return
	}

	if !newTransaction.IsValid() {
		e := apierror.NewBadRequest(fmt.Sprintf("invalid transaction type: %s", newTransaction.GetType()))
		c.JSON(
			e.Status(),
			e.JSON(),
		)
		return
	}

	response, err := t.p.NewTransaction(newTransaction)
	if err != nil {
		if err == errorx.NotEnoughFunds {
			e := apierror.NewUnprocessableEntity(fmt.Sprintf("error processing transaction: %s", err.Error()))
			c.JSON(
				e.Status(),
				e.JSON(),
			)
			return
		}
		e := apierror.NewInternalServerError(fmt.Sprintf("error processing transaction: %s", err.Error()))
		c.JSON(
			e.Status(),
			e.JSON(),
		)
		return
	}

	c.JSON(http.StatusCreated, response)
	return
}

func (t *transaction) HandleFindTransactionByID(c *gin.Context) {

	id := c.Param("id")

	if id == "" {
		e := apierror.NewUnprocessableEntity("invalid id supplied")
		c.JSON(
			e.Status(),
			e.JSON(),
		)
		return
	}

	response, err := t.p.FindTransactionByID(id)

	if err != nil {
		var e apierror.Error
		if errors.Is(err, errorx.NotFound) {
			e = apierror.NewNotFoundError(fmt.Sprintf("error finding transaction: %s", err.Error()))
			c.JSON(
				e.Status(),
				e.JSON(),
			)
			return
		}
		e = apierror.NewInternalServerError(fmt.Sprintf("error finding transaction: %s", err.Error()))
		c.JSON(
			e.Status(),
			e.JSON(),
		)
		return
	}

	c.JSON(http.StatusOK, response)
	return
}

func (t *transaction) HandleFetchAllTransactions(c *gin.Context) {

	response, err := t.p.ListAllTransactions()
	if err != nil {
		e := apierror.NewInternalServerError(fmt.Sprintf("error fetching all transactions: %s", err.Error()))
		c.JSON(
			e.Status(),
			e.JSON(),
		)
		return
	}
	c.JSON(http.StatusOK, response)
	return
}

//NewTransactionHandler returns a
func NewTransactionHandler() Transaction {
	return &transaction{
		p: processor.NewTransactionProcessor(),
	}
}
