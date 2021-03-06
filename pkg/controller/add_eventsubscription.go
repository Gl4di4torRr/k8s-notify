package controller

import (
	"github.com/redhat-cop/events-notifier/pkg/controller/eventsubscription"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, eventsubscription.Add)
}
