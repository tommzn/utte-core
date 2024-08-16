package core

import (
	uuid "github.com/satori/go.uuid"
	model "github.com/tommzn/utte-model"
)

func NewIdentifier() model.Identifier {
	return model.Identifier(uuid.NewV4().String())
}
