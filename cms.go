package main

import (
	EComApp "github.com/codedv8/go-ecom-app"
)

// CMS - The main struct for this module
type CMS struct {
	App     *EComApp.Application
	Message string
}

var cmsList []CMS

// Exports

// SysInit - Pre initialization of this module
func SysInit(app *EComApp.Application) error {
	cms := &CMS{
		App:     app,
		Message: "Welcome to the CMS plugin",
	}
	cms.SysInit(app)

	cmsList = append(cmsList, *cms)

	return nil
}

// Init - Initialization of this module
func Init(app *EComApp.Application) error {
	for _, cms := range cmsList {
		cms.Init(app)
	}

	return nil
}

// Done - Shut down of this module
func Done(app *EComApp.Application) error {
	for _, cms := range cmsList {
		cms.Done(app)
	}

	return nil
}

var cms CMS
