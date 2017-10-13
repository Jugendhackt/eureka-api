package queries

import (
	"github.com/graphql-go/graphql"
	"../../models"
	gqlModels "../models"
)

var User = graphql.Field{
	Type: gqlModels.User,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		return models.User{"A", "A", "A", 5},  nil
	},
}