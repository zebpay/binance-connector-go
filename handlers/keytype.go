package handlers

import (
	"crypto/ed25519"
	"crypto/x509"
	"encoding/pem"
	"log"
)

func ParseEd25519PEM(pemKey string) ed25519.PrivateKey {
	block, _ := pem.Decode([]byte(pemKey))
	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		log.Fatal(err)
	}
	return key.(ed25519.PrivateKey)
}
