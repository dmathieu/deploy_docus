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

func (r *Repository) Token() string {
	value, _ := r.Rsa.Encrypt([]byte(r.Name()))
	return value
}

func (r *Repository) IsNew() bool {
	return r.Id == 0
}

func (r *Repository) Save() error {
	var id int64

	if r.IsNew() {
		row, err := QueryRow(`INSERT INTO repositories (origin, destination, rsa_key) VALUES ($1, $2, $3) RETURNING id;`, r.Origin, r.Destination, r.Rsa.Key)
		if err != nil {
			return err
		}
		err = row.Scan(&id)
		if err != nil {
			return err
		}
		r.Id = id
	} else {
		_, err := QueryRow(`UPDATE repositories SET origin = $1, destination = $2, rsa_key = $3 WHERE id = $4`, r.Origin, r.Destination, r.Rsa.Key, r.Id)
		if err != nil {
			return err
		}
	}

	return nil
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

func FindRepository(id int64) (*Repository, error) {
	var (
		origin      string
		destination string
		rsa_key     []byte
	)
	row, err := QueryRow(`SELECT origin, destination, rsa_key FROM repositories WHERE id = $1`, id)
	if err != nil {
		return nil, err
	}
	err = row.Scan(&origin, &destination, &rsa_key)
	if err != nil {
		return nil, err
	}

	return BuildRepository(id, origin, destination, rsa_key), nil
}

func AllRepositories() ([]Repository, error) {
	var count int64
	row, err := QueryRow(`SELECT COUNT(id) FROM repositories`)
	if err != nil {
		return nil, err
	}
	row.Scan(&count)

	repositories := make([]Repository, count)
	rows, err := Query(`SELECT id, origin, destination, rsa_key FROM repositories`)
	if err != nil {
		return nil, err
	}

	i := 0
	for rows.Next() {
		var (
			id          int64
			origin      string
			destination string
			rsa_key     []byte
		)

		rows.Scan(&id, &origin, &destination, &rsa_key)
		repository := BuildRepository(id, origin, destination, rsa_key)
		repositories[i] = *repository
		i += 1
	}

	return repositories, nil
}
