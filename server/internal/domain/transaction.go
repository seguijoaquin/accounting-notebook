package domain

const (
	//CREDIT refers to a credit transaction type
	CREDIT = "credit"
	//DEBIT refers to a credit transaction type
	DEBIT = "debit"
)

//TransactionType represents any type of transaction [debit, credit]
type TransactionType string

//Transaction represents a transaction of a money account
type Transaction struct {
	ID string `json:"id"`
	TransactionBody
	EffectiveDate string `json:"effective_date"`
}

//TransactionBody represents a transaction request made to a money account
type TransactionBody struct {
	Type    TransactionType `json:"type" binding:"required"`
	Ammount int             `json:"ammount" binding:"required"`
}

//GetAmmount returns a transactions ammount
func (tb *TransactionBody) GetAmmount() int {
	return tb.Ammount
}

//GetType returns a transactions type
func (tb *TransactionBody) GetType() TransactionType {
	return tb.Type
}

//IsValid verifies if a transaction body request is valid or not
func (tb *TransactionBody) IsValid() bool {
	return (tb.Type == CREDIT || tb.Type == DEBIT)
}

//IsDebit verifies if the transaction is of type DEBIT
func (tb *TransactionBody) IsDebit() bool {
	return tb.Type == DEBIT
}

//IsCredit verifies if the transaction is of type CREDIT
func (tb *TransactionBody) IsCredit() bool {
	return tb.Type == CREDIT
}
