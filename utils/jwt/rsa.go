package jwt

import (
	"crypto/rsa"
	"time"

	j "github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type r struct{}

func (r *r) Token(UserId string, key *rsa.PrivateKey, duration time.Duration) (string, error) {
	t := j.New(j.GetSigningMethod("RS256"))

	t.Claims = &j.StandardClaims{
		ExpiresAt: time.Now().Add(duration).Unix(),
		Id:        uuid.NewString(),
		IssuedAt:  time.Now().Unix(),
		Subject:   UserId,
	}

	return t.SignedString(key)
}

func (r *r) Verify(token string, key *rsa.PublicKey) (*j.StandardClaims, error) {
	t, e := j.ParseWithClaims(token, &j.StandardClaims{}, func(t *j.Token) (interface{}, error) {
		if _, ok := t.Method.(*j.SigningMethodRSA); !ok {
			return nil, ErrUnexpectedSigningMethod
		}
		return key, nil
	})

	if e != nil {
		return nil, e
	}

	c, ok := t.Claims.(*j.StandardClaims)
	if !ok {
		return nil, ErrUnableToMarshalClaim
	}

	return c, nil
}
