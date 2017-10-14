package queries

import (
	"github.com/graphql-go/graphql"
)

var Root = graphql.NewObject(graphql.ObjectConfig{
	Name: "rootQuery",
	Fields: graphql.Fields{
		"user": &User,

		//"event": &Event,
	},
})