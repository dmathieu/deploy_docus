package deploy_docus

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Rsa struct {
	Repository *Repository
	Private    *rsa.PrivateKey
	Key        []byte
}

func (r *Rsa) Encrypt(value []byte) ([]byte, error) {
	var err error

	var out []byte
	out, err = rsa.EncryptOAEP(sha1.New(), rand.Reader, &r.Private.PublicKey, value, nil)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func (r *Rsa) Decrypt(value []byte) ([]byte, error) {
	var err error

	var out []byte
	out, err = rsa.DecryptOAEP(sha1.New(), rand.Reader, r.Private, value, nil)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func (r *Rsa) WriteKey() {
	path := r.KeyPath()

	err := os.MkdirAll(filepath.Dir(path), 0700)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(path, r.Key, 0700)
	if err != nil {
		panic(err)
	}
}

func (r *Rsa) KeyPath() string {
	return fmt.Sprintf("/tmp/deploy_docus/keys/%s", r.Repository.Name())
}

func BuildPrivateKey(content []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(content)
	return x509.ParsePKCS1PrivateKey(block.Bytes)
}

func NewRsa(repository *Repository) *Rsa {
	rsa := &Rsa{Repository: repository, Key: []byte(repository.PKey)}

	key, err := BuildPrivateKey(rsa.Key)
	if err != nil {
		panic(err)
	}
	rsa.Private = key
	rsa.WriteKey()

	return rsa
}
