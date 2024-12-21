package calculate

import "errors"

var (
	ErrInvalidExpression = errors.New("invalid expression")
	ErrInternalServer    = errors.New("internal server error")
)
