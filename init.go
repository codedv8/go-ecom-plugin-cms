package main

import (
	"context"
	"log"

	EComApp "github.com/codedv8/go-ecom-app"
	EComStructs "github.com/codedv8/go-ecom-structs"
	_ "github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// Page - Basic definition af a page
type Page struct {
	Title  string `bson:"title"`
	SEOURL string `bson:"seourl"`
}

// SysInit - Pre initialization of the object
func (cms *CMS) SysInit(app *EComApp.Application) {

}

// Init - Initialization of the object
func (cms *CMS) Init(app *EComApp.Application) {
	app.ListenToHook("ROUTER_WILDCARD", func(payload interface{}) (bool, error) {
		switch c := payload.(type) {
		case *EComStructs.RouterWildcard:
			path := c.Context.Request.URL.Path
			if path == "/" {
				collection := app.DB.Client.Database("shop").Collection("cms")
				var page bson.M
				err := collection.FindOne(context.TODO(), bson.D{}).Decode(&page)
				if err != nil {
					log.Printf("Error FindOne %+v", err)
					return false, err
				}
				c.Context.JSON(200, page)
				return false, nil
			}
			return true, nil
		}
		return true, nil
	})
}
