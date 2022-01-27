package jwt

import (
	"time"

	j "github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type h struct{}

func (h *h) Token(userId string, key string, duration time.Duration) (string, error) {
	t := j.New(j.GetSigningMethod("HS256"))

	t.Claims = &j.StandardClaims{
		ExpiresAt: time.Now().Add(duration).Unix(),
		Id:        uuid.NewString(),
		IssuedAt:  time.Now().Unix(),
		Subject:   userId,
	}

	return t.SignedString([]byte(key))
}

func (h *h) Verify(token string, key string) (*j.StandardClaims, error) {
	t, e := j.ParseWithClaims(token, &j.StandardClaims{}, func(t *j.Token) (interface{}, error) {
		if _, ok := t.Method.(*j.SigningMethodHMAC); !ok {
			return nil, ErrUnexpectedSigningMethod
		}
		return []byte(key), nil
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
