package er

import "errors"

var (
	ErrNoRows       = errors.New("no rows in result set")
  ErrInvalidType  = errors.New("invalid symbols in type")
  ErrInvalidLimit = errors.New("limit is too high")
)
