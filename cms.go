package main

import (
	ecomapp "github.com/codedv8/go-ecom-app"
)

// CMS - The main struct for this module
type CMS struct {
	App     *ecomapp.Application
	Message string
}

var cmsList []CMS

// Exports

// SysInit - Pre initialization of this module
func SysInit(app *ecomapp.Application) error {
	cms := &CMS{
		App:     app,
		Message: "Welcome to the CMS plugin",
	}
	cms.SysInit(app)

	cmsList = append(cmsList, *cms)

	return nil
}

// Init - Initialization of this module
func Init(app *ecomapp.Application) error {
	for _, cms := range cmsList {
		cms.Init(app)
	}

	return nil
}

// Done - Shut down of this module
func Done(app *ecomapp.Application) error {
	for _, cms := range cmsList {
		cms.Done(app)
	}

	return nil
}

var cms CMS
