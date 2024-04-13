package crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"

	"github.com/pinkphantasm/hieda/src/sign_service/internal/pkg/config"
)

type Adapter struct {
	privateKey *rsa.PrivateKey
}

func NewAdapter() *Adapter {
	cfg := config.New()

	privateKeyBytes, err := os.ReadFile(cfg.PrivateKey)
	if err != nil {
		panic(err)
	}

	privateKeyBlock, _ := pem.Decode(privateKeyBytes)

	privateKey, err := x509.ParsePKCS8PrivateKey(privateKeyBlock.Bytes)
	if err != nil {
		panic(err)
	}

	return &Adapter{
		privateKey: privateKey.(*rsa.PrivateKey),
	}
}

func (a Adapter) Sign(data []byte) ([]byte, error) {
	return rsa.EncryptPKCS1v15(rand.Reader, &a.privateKey.PublicKey, data)
}

func (a Adapter) Verify(data []byte) ([]byte, error) {
	return rsa.DecryptPKCS1v15(rand.Reader, a.privateKey, data)
}
