package processor

import (
	"github.com/seguijoaquin/accounting-notebook/internal/repository"
)

//Account represents an object that can process money account requests
type Account interface {
	GetBalance() (int, error)
}

type account struct {
}

//GetBalance returns the total balance of the account
func (ac *account) GetBalance() (int, error) {
	transactions, err := repository.GetInstance().ListAll()
	if err != nil {
		return 0, err
	}

	balance := 0

	for _, t := range transactions {
		if t.IsDebit() {
			balance -= t.GetAmmount()
		}
		if t.IsCredit() {
			balance += t.GetAmmount()
		}
	}

	return balance, nil
}

//NewAccountProcessor returns a processor object that can operate accounts
func NewAccountProcessor() Account {
	return &account{}
}
