package controller

import (
	"github.com/platform9/alertmgr-config/pkg/controller/alertmgrcfg"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, alertmgrcfg.Add)
}
