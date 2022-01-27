package jwt

import (
	"errors"
)

var (
	ErrUnableToMarshalClaim    = errors.New("unable to marshal token claim")
	ErrUnexpectedSigningMethod = errors.New("unexpected signing method")
)

var (
	RSA  = &r{}
	HMAC = &h{}
)
