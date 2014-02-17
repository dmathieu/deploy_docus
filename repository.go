package deploy_docus

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type Repository struct {
	Origin      string
	Destination string
	PKey        string
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

func (r *Repository) CreatePKey() {
	path := r.PKeyPath()

	err := os.MkdirAll(filepath.Dir(path), 0700)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(path, []byte(r.PKey), 0700)
	if err != nil {
		panic(err)
	}
}

func (r *Repository) PKeyPath() string {
	return fmt.Sprintf("/tmp/deploy_docus/keys/%s", r.Name())
}

func FindRepository() *Repository {
	repository := &Repository{
		Origin:      os.Getenv("REPOSITORY_ORIGIN"),
		Destination: os.Getenv("REPOSITORY_DESTINATION"),
		PKey:        os.Getenv("REPOSITORY_PKEY"),
	}
	repository.CreatePKey()

	return repository
}
