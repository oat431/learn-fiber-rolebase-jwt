package bootstrap

import (
	"github.com/jmoiron/sqlx"
)

type APIContainer struct {
}

func NewAPIContainer(db *sqlx.DB) *APIContainer {

	return &APIContainer{}
}
