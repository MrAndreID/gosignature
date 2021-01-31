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

type Unsigner interface {
	Unsign(data []byte, signature []byte) error
}

type rsaPrivateKey struct {
	*rsa.PrivateKey
}

type rsaPublicKey struct {
	*rsa.PublicKey
}

func (rsaPrivateKey *rsaPrivateKey) Sign(data []byte) ([]byte, error) {
	hash := sha256.New()
	hash.Write(data)

	hashed := hash.Sum(nil)

	return rsa.SignPKCS1v15(rand.Reader, rsaPrivateKey.PrivateKey, crypto.SHA256, hashed)
}

func (rsaPublicKey *rsaPublicKey) Unsign(data []byte, signature []byte) error {
	hash := sha256.New()
	hash.Write(data)

	hashed := hash.Sum(nil)

	return rsa.VerifyPKCS1v15(rsaPublicKey.PublicKey, crypto.SHA256, hashed, signature)
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

func Verify(signature string, publicKey []byte, value []byte) (bool, error) {
	parser, err := parsePublicKey(publicKey)

	if err != nil {
		return false, errors.New("failed to parse the public key.")
	}

	key, err := base64.StdEncoding.DecodeString(signature)

	if err != nil {
		return false, errors.New("invalid signature.")
	}
	
	if parser.Unsign(value, key) != nil {
		return false, errors.New("could not sign the request.")
	}

	return true, nil
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
		return nil, errors.New("key type not supported [4042].")
	}

	return sshKey, nil
}

func parsePublicKey(publicKey []byte) (Unsigner, error) {
	block, _ := pem.Decode(publicKey)

	if block == nil {
		return nil, errors.New("no key was found.")
	}

	var rawPublicKey interface{}

	switch block.Type {
	case "PUBLIC KEY":
		rsa, err := x509.ParsePKIXPublicKey(block.Bytes)

		if err != nil {
			return nil, err
		}

		rawPublicKey = rsa
	default:
		return nil, errors.New("key type not supported [4041].")
	}

	return newUnsignerFromPublicKey(rawPublicKey)
}

func newUnsignerFromPublicKey(rawPublicKey interface{}) (Unsigner, error) {
	var sshKey Unsigner

	switch publicKeyType := rawPublicKey.(type) {
	case *rsa.PublicKey:
		sshKey = &rsaPublicKey{publicKeyType}
	default:
		return nil, errors.New("key type not supported [4042].")
	}

	return sshKey, nil
}
