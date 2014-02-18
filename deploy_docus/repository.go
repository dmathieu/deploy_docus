package deploy_docus

import (
	"fmt"
	"os"
	"strings"
)

type Repository struct {
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

func FindRepository() *Repository {
	key := []byte(os.Getenv("REPOSITORY_PKEY"))
	repository := &Repository{
		Origin:      os.Getenv("REPOSITORY_ORIGIN"),
		Destination: os.Getenv("REPOSITORY_DESTINATION"),
	}
	repository.Rsa = NewRsa(repository, key)

	return repository
}
