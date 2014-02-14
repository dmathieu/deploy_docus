package deploy_docus

import (
	"os"
)

type Repository struct {
	Origin      string
	Destination string
}

func FindRepository() *Repository {
	return &Repository{
		Origin:      os.Getenv("REPOSITORY_ORIGIN"),
		Destination: os.Getenv("REPOSITORY_DESTINATION"),
	}
}
