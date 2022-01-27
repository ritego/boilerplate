package sec

import (
	"crypto/md5"
	cryptoRand "crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	mathRand "math/rand"
	"os"
	"strconv"
	"time"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandString(n int64) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[mathRand.Intn(len(letters))]
	}
	return string(b)
}

func Keys() (*rsa.PrivateKey, *rsa.PublicKey, string, string) {
	private, _ := rsa.GenerateKey(cryptoRand.Reader, 2048)
	public := &private.PublicKey

	privateBytes := pem.EncodeToMemory(
		&pem.Block{
			Type:  "PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(private),
		},
	)

	pubASN1, _ := x509.MarshalPKIXPublicKey(public)
	publicBytes := pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pubASN1,
	})

	return private, public, string(privateBytes), string(publicBytes)
}

var (
	hostName, _ = os.Hostname()
	hostHex     = md5.Sum([]byte(hostName))
	host        = hex.EncodeToString(hostHex[:])[:5] // 6 xters
	processId   = strconv.Itoa(os.Getegid())[:2]     // 3 xters

)

func GenerateReference() string {
	return fmt.Sprintf("%d-%s-%s-%s-%s", time.Now().Unix(), host, "JRA", processId, RandString(7))[:32]
}
