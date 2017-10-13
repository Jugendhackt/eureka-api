package queries

import (
	"github.com/graphql-go/graphql"
	"../../models"
	gqlModels "../models"
	"time"
)

var User = graphql.Field{
	Type: gqlModels.User,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		return models.User{time.Now(), "A", "A", "A", 500},  nil
	},
}