package deploy_docus

import (
	"fmt"
	"strings"
)

type Repository struct {
	Id          int64
	Origin      string
	Destination string

	Rsa *Rsa
}

func (r *Repository) Name() string {
	suffix := "git@github.com:"
	affix := ".git"
	name := r.Origin[len(suffix) : len(r.Origin)-len(affix)]

	return strings.Replace(name, "/", "_", 1)
}

func (r *Repository) LocalPath() string {
	return fmt.Sprintf("/tmp/%s", r.Name())
}

func BuildRepository(id int64, origin string, destination string, rsa_key []byte) *Repository {
	repository := &Repository{
		Id:          id,
		Origin:      origin,
		Destination: destination,
	}
	repository.Rsa = NewRsa(repository, rsa_key)
	return repository
}

func CreateRepository(origin string, destination string, rsa_key []byte) (*Repository, error) {
	db, err := GetDbConnection()
	if err != nil {
		return nil, err
	}

	var id int64
	err = db.QueryRow(`INSERT INTO repositories (origin, destination, rsa_key) VALUES ($1, $2, $3) RETURNING id;`, origin, destination, rsa_key).Scan(&id)
	if err != nil {
		return nil, err
	}

	return BuildRepository(id, origin, destination, rsa_key), nil
}

func FindRepository(id int64) (*Repository, error) {
	db, err := GetDbConnection()
	if err != nil {
		return nil, err
	}

	var origin string
	var destination string
	var rsa_key []byte
	err = db.QueryRow("SELECT origin, destination, rsa_key FROM repositories WHERE id = $1", id).Scan(&origin, &destination, &rsa_key)
	if err != nil {
		return nil, err
	}

	return BuildRepository(id, origin, destination, rsa_key), nil
}
