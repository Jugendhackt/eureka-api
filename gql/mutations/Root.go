package mutations

import (
	"github.com/graphql-go/graphql"
)

var Root = graphql.NewObject(graphql.ObjectConfig{
	Name: "rootMutation",
	Fields: graphql.Fields{
		"createUser": &CreateUser,
		"updateUser": &UpdateUser,

		"createEvent": &CreateEvent,
	},
})