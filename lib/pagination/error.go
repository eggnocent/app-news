package pagination

import "errors"

var (
	ErrorHexPage     = errors.New("cooshen page more than total page")
	ErrorPage        = errors.New("page must greater than 0")
	ErrorPageEmpty   = errors.New("page cannot be empty")
	ErrorPageInvalid = errors.New("page invalid, must be number")
)
