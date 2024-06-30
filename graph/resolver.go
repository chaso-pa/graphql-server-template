package graph

import (
	"github/chaso-pa/gql-server/graph/model"

	"gorm.io/gorm"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DB    *gorm.DB
	todos []*model.Todo
	users []*model.User
}
