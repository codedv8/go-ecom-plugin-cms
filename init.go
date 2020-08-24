package main

import (
	EComApp "github.com/codedv8/go-ecom-app"
	EComStructs "github.com/codedv8/go-ecom-structs"
)

func (cms *CMS) SysInit(app *EComApp.Application) {

}

func (cms *CMS) Init(app *EComApp.Application) {
	app.ListenToHook("ROUTER_WILDCARD", func(payload interface{}) (bool, error) {
		switch c := payload.(type) {
		case *EComStructs.RouterWildcard:
			path := c.Context.Request.URL.Path
			if len(path) == 10 {
				c.Context.String(200, "Path was 10 chars in length")
				return false, nil
			}
			return true, nil
		}
		return true, nil
	})
}
