package error

import "errors"

var(
	ErrDBConnect = errors.New("DB connection fail")
)