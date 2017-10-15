package models

import (
	"github.com/graphql-go/graphql"
)

var EventResults = graphql.NewObject(graphql.ObjectConfig{
	Name: "EventResult",
	Fields: graphql.Fields{
		"recent": &graphql.Field{
			Type: graphql.NewList(Event),
		},
		"upcoming": &graphql.Field{
			Type: graphql.NewList(Event),
		},
	},
})