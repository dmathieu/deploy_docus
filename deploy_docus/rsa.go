package deploy_docus

import (
	"code.google.com/p/go.crypto/ssh"
	"crypto/rand"
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
}

func (r *Rsa) Encrypt(value []byte) (string, error) {
	hasher := sha1.New()
	key, err := r.PrivateKey()
	if err != nil {
		return "", err
	}

	hasher.Write([]byte(fmt.Sprintf("%s-%s", key, value)))
	return base64.URLEncoding.EncodeToString(hasher.Sum(nil)), nil
}

func (r *Rsa) WriteKey() error {
	path := r.KeyPath()

	err := os.MkdirAll(filepath.Dir(path), 0700)
	if err != nil {
		return err
	}

	key, _ := r.PrivateKey()
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(path, key, 0600)
	return err
}

func (r *Rsa) PrivateKey() ([]byte, error) {
	cert := x509.MarshalPKCS1PrivateKey(r.Private)
	blk := new(pem.Block)
	blk.Type = "RSA PRIVATE KEY"
	blk.Bytes = cert
	content := pem.EncodeToMemory(blk)

	return content, nil
}

func (r *Rsa) PublicKey() string {
	key, _ := ssh.NewPublicKey(&r.Private.PublicKey)
	marshalled := ssh.MarshalPublicKey(key)
	encoded := base64.StdEncoding.EncodeToString(marshalled) + "\n"

	return fmt.Sprintf("ssh-rsa %s", encoded)
}

func (r *Rsa) KeyPath() string {
	return fmt.Sprintf("/tmp/deploy_docus/keys/%s", r.Repository.Name())
}

func BuildPrivateKey(content []byte) (*rsa.PrivateKey, error) {
	if content == nil {
		return rsa.GenerateKey(rand.Reader, 1024)
	} else {
		block, _ := pem.Decode(content)
		return x509.ParsePKCS1PrivateKey(block.Bytes)
	}
}

func BuildRsa(repository *Repository, ssh_key []byte) *Rsa {
	key, err := BuildPrivateKey(ssh_key)
	if err != nil {
		panic(err)
	}
	rsa := &Rsa{repository, key}

	return rsa
}
