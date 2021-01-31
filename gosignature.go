package gosignature

import (
	"errors"
	"crypto"
	"crypto/rsa"
	"crypto/x509"
	"crypto/rand"
	"encoding/pem"
	"crypto/sha256"
	"encoding/base64"
)

type Signer interface {
	Sign(data []byte) ([]byte, error)
}

type rsaPrivateKey struct {
	*rsa.PrivateKey
}

func (rsaPrivateKey *rsaPrivateKey) Sign(data []byte) ([]byte, error) {
	hash := sha256.New()
	hash.Write(data)

	hashed := hash.Sum(nil)

	return rsa.SignPKCS1v15(rand.Reader, rsaPrivateKey.PrivateKey, crypto.SHA256, hashed)
}

func Generate(privateKey []byte, value []byte) (string, error) {
	parser, err := parsePrivateKey(privateKey)

	if err != nil {
		return "", errors.New("failed to parse the private key.")
	}

	signed, err := parser.Sign(value)

	if err != nil {
		return "", errors.New("could not sign the request.")
	}

	return base64.StdEncoding.EncodeToString(signed), nil
}

func parsePrivateKey(privateKey []byte) (Signer, error) {
	block, _ := pem.Decode(privateKey)

	if block == nil {
		return nil, errors.New("no key was found.")
	}

	var rawPrivateKey interface{}

	switch block.Type {
	case "RSA PRIVATE KEY":
		rsa, err := x509.ParsePKCS1PrivateKey(block.Bytes)

		if err != nil {
			return nil, err
		}

		rawPrivateKey = rsa
	default:
		return nil, errors.New("key type not supported [4041].")
	}

	return newSignerFromPrivateKey(rawPrivateKey)
}

func newSignerFromPrivateKey(rawPrivateKey interface{}) (Signer, error) {
	var sshKey Signer

	switch privateKeyType := rawPrivateKey.(type) {
	case *rsa.PrivateKey:
		sshKey = &rsaPrivateKey{privateKeyType}
	default:
		return nil, errors.New("key type not supported [4041].")
	}

	return sshKey, nil
}
