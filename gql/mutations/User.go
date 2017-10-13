package mutations

import (
	"github.com/graphql-go/graphql"
	gqlModels "../models"
	"../../models"
	"fmt"
	"time"
)

var CreateUser = graphql.Field{
	Type: gqlModels.User,
	Args: graphql.FieldConfigArgument{
		"emailAddress": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {

		// TODO: Validate if email is correct
		mail, isOk := params.Args["emailAddress"].(string)

		if !isOk { return nil, fmt.Errorf("mail couldn't be turned into a string")}
		user := models.User{JoinDate: time.Now(), EmailAddress: mail,  ExperiencePoints: 500}

		// TODO: Insert in database.

		return user, nil
	},
}