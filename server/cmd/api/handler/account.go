package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/seguijoaquin/accounting-notebook/cmd/api/apierror"
	"github.com/seguijoaquin/accounting-notebook/internal/domain"
	"github.com/seguijoaquin/accounting-notebook/internal/processor"
)

//Account represents an money account object
type Account interface {
	HandleBalance(c *gin.Context)
}

type accout struct {
	p processor.Account
}

//HandleBalance receives balance requests and is responsible for replying them
func (ac *accout) HandleBalance(c *gin.Context) {

	response, err := ac.p.GetBalance()
	if err != nil {
		e := apierror.NewInternalServerError(fmt.Sprintf("error getting balance: %s", err.Error()))
		c.JSON(
			e.Status(),
			e.JSON(),
		)
		return
	}
	c.JSON(
		http.StatusOK,
		domain.Account{
			Balance: response,
		},
	)
	return
}

//NewAccountHandler returns a suitable handler for managing money accounts
func NewAccountHandler() Account {
	return &accout{
		p: processor.NewAccountProcessor(),
	}
}
