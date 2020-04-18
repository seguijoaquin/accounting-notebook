package repository

import (
	"strconv"
	"sync"
	"time"

	"github.com/seguijoaquin/accounting-notebook/internal/domain"
	"github.com/seguijoaquin/accounting-notebook/internal/errorx"
)

var (
	r    *repository // The one and only repository
	once sync.Once
)

//GetInstance returns a Singleton instance of a repository or database
func GetInstance() *repository {
	once.Do(func() {
		r = &repository{
			transactions: make(map[string]domain.Transaction),
		}
	})

	return r
}

type repository struct {
	transactions map[string]domain.Transaction
	indexCount   int
	mu           sync.RWMutex // To sync operations at db layer
}

//Save persists a new transaction into the storage and returns the persisted object
func (r *repository) Save(data domain.Transaction) (domain.Transaction, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.indexCount++
	id := strconv.Itoa(r.indexCount)
	data.ID = id
	data.EffectiveDate = time.Now().Format(time.RFC3339)

	r.transactions[id] = data

	return data, nil
}

//Get returns a transaction given it's id or an error
func (r *repository) Get(id string) (domain.Transaction, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	t, ok := r.transactions[id]
	if !ok {
		return domain.Transaction{}, errorx.NotFound
	}
	return t, nil
}

//ListAll fetches all transactions and returns them unordered
func (r *repository) ListAll() ([]domain.Transaction, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	transactions := []domain.Transaction{}

	for _, t := range r.transactions {
		transactions = append(transactions, t)
	}

	return transactions, nil
}
