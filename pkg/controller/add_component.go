package controller

import (
	"github.com/redhat-developer/devconsole-operator/pkg/controller/component"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, component.Add)
}
