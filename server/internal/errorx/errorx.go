package errorx

type customError string

func (ce customError) Error() string {
	return string(ce)
}

//NotFound represents a NotFoundError
const NotFound customError = "not found"

//AlreadyExists represents a AlreadyExistsError
const AlreadyExists customError = "already exists"

//UnprocesableEntity represents a UnprocesableEntityError
const UnprocesableEntity customError = "unprocessable entity"

//NotEnoughFunds represents a situation when a debit transaction exceeds your funds
const NotEnoughFunds customError = "not enough funds"
