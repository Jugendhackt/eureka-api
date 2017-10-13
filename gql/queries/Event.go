package queries

import (
	"github.com/graphql-go/graphql"
	gqlModels "../models"
	"../../models"
	"time"
)

var Event = graphql.Field{
	Type: gqlModels.Event,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		return models.Event{"hello", "FD", time.Now(), "DE"}, nil
	},
}