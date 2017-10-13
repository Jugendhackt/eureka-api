package models

import (
	"github.com/graphql-go/graphql"
)

var Event = graphql.NewObject(graphql.ObjectConfig{
	Name: "Event",
	Fields: graphql.Fields{
		"creationDate": &graphql.Field{
			Type: graphql.DateTime,
		},
		"creatorId": &graphql.Field{
			Type: graphql.String,
		},
		"place": &graphql.Field{
			Type: graphql.String,
		},
		"lang": &graphql.Field{
			Type: graphql.String,
		},
		"time": &graphql.Field{
			Type: graphql.DateTime,
		},
		"activityId": &graphql.Field{
			Type: graphql.String, // Id
		},
	},
})