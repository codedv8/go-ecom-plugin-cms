package main

import (
	EComApp "github.com/codedv8/go-ecom-app"
)

func (cms *CMS) SysInit(app *EComApp.Application) {

}

func (cms *CMS) Init(app *EComApp.Application) {
	//app.ListenToHook("ROUTER_WILDCARD", func(payload interface{}) (bool, error) {
	//    switch v := payload.(type) {
	//    case *EComStructsAPI.RouterWildcard:
	//        log.Printf("ROUTER_WILDCARD: %+v\n", v)
	//        v.S = "Bye bye"
	//    }
	//    return true, nil
	//})
}
