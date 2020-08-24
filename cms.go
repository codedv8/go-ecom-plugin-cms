package main

import (
	EComApp "github.com/codedv8/go-ecom-app"
)

type CMS struct {
	App     *EComApp.Application
	Message string
}

var cmsList []CMS

// Exports
func SysInit(app *EComApp.Application) error {
	cms := &CMS{
		App:     app,
		Message: "Welcome to the API plugin",
	}
	cms.SysInit(app)

	cmsList = append(cmsList, *cms)

	return nil
}

func Init(app *EComApp.Application) error {
	for _, cms := range cmsList {
		cms.Init(app)
	}

	return nil
}

func Done(app *EComApp.Application) error {
	for _, cms := range cmsList {
		cms.Done(app)
	}

	return nil
}

var cms CMS
