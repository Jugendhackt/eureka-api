package models

import (
	"github.com/graphql-go/graphql"
)

var Event = graphql.NewObject(graphql.ObjectConfig{
	Name: "Event",
	Fields: graphql.Fields{
		"place": &graphql.Field{
			Type: graphql.String,
		},
		"lang": &graphql.Field{
			Type: graphql.String,
		},
		"time": &graphql.Field{
			Type: graphql.DateTime,
		},
		"activity": &graphql.Field{
			Type: graphql.String, // Id
		},
	},
})