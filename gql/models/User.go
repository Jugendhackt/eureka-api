package models

import (
	"github.com/graphql-go/graphql"
)

var User = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"joinDate": &graphql.Field{
			Type: graphql.DateTime,
		},
		"emailAddress": &graphql.Field{
			Type: graphql.String,
		},
		"password": &graphql.Field{
			Type: graphql.String,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"hometown": &graphql.Field{
			Type: graphql.String,
		},
		"experiencePoints": &graphql.Field{
			Type: graphql.Int,
		},
	},
})