package processor

import (
	"sync"

	"github.com/seguijoaquin/accounting-notebook/internal/domain"
	"github.com/seguijoaquin/accounting-notebook/internal/errorx"
	"github.com/seguijoaquin/accounting-notebook/internal/repository"
)

//Transaction represents an object that can process transactions
type Transaction interface {
	NewTransaction(t domain.TransactionBody) (*domain.Transaction, error)
	FindTransactionByID(id string) (*domain.Transaction, error)
	ListAllTransactions() ([]domain.Transaction, error)
}

type transaction struct {
	ac Account
	mu sync.RWMutex // Needed to simulate read and store transaction in NewTransaction
}

//NewTransaction sends a new transaction to the storage to be saved
func (t *transaction) NewTransaction(tb domain.TransactionBody) (*domain.Transaction, error) {
	t.mu.Lock()
	defer t.mu.Unlock()

	if tb.IsDebit() && !t.hasEnoughFunds(tb) {
		return nil, errorx.NotEnoughFunds
	}

	var transaction domain.Transaction
	transaction.Ammount = tb.GetAmmount()
	transaction.Type = tb.GetType()

	transaction, err := repository.GetInstance().Save(transaction)
	if err != nil {
		return nil, err
	}

	return &transaction, nil
}

func (t *transaction) hasEnoughFunds(tb domain.TransactionBody) bool {
	b, err := t.ac.GetBalance()
	if err != nil {
		return false
	}
	return ((b - tb.Ammount) >= 0)
}

//FindTransactionByID searchs for an specific transaction in the storage
func (t *transaction) FindTransactionByID(id string) (*domain.Transaction, error) {
	t.mu.RLock()
	defer t.mu.RUnlock()

	transaction, err := repository.GetInstance().Get(id)
	if err != nil {
		return nil, err
	}

	return &transaction, nil
}

//ListAllTransactions fetches all transactions from the storage and returns them unordered
func (t *transaction) ListAllTransactions() ([]domain.Transaction, error) {
	t.mu.RLock()
	defer t.mu.RUnlock()

	transactions, err := repository.GetInstance().ListAll()
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

//NewTransactionProcessor returns a processor object that can CRUD transactions
func NewTransactionProcessor() Transaction {
	return &transaction{
		ac: NewAccountProcessor(),
	}
}
