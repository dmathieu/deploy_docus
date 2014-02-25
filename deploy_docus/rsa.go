package deploy_docus

import (
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
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

func (r *Rsa) Encrypt(value []byte) (string, error) {
	hasher := sha1.New()
	hasher.Write([]byte(fmt.Sprintf("%s-%s", r.Key, value)))
	return base64.URLEncoding.EncodeToString(hasher.Sum(nil)), nil
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

func NewRsa(repository *Repository, ssh_key []byte) *Rsa {
	rsa := &Rsa{Repository: repository, Key: ssh_key}

	key, err := BuildPrivateKey(rsa.Key)
	if err != nil {
		panic(err)
	}
	rsa.Private = key
	rsa.WriteKey()

	return rsa
}
