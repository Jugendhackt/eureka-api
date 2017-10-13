package mutations

import (
	"github.com/graphql-go/graphql"
	gqlModels "../models"
	"fmt"
	"time"
	"../../models"
)

var CreateEvent = graphql.Field{
	Type: gqlModels.Event,
	Args: graphql.FieldConfigArgument{
			"creatorId": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"place": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"activityId": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"time": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"lang": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		creatorId, isOk := params.Args["creatorId"].(string)
		if !isOk { return nil, getErr("creatorId", "string") }

		place, isOk := params.Args["place"].(string)
		if !isOk { return nil, getErr("place", "string") }

		activityId, isOk := params.Args["activityId"].(string)
		if !isOk { return nil, getErr("activityId", "string") }

		aTime, isOk := params.Args["time"].(string)
		if !isOk { return nil, getErr("time", "string") }

		lang, isOk := params.Args["lang"].(string)
		if !isOk { return nil, getErr("lang", "string") }

		parsedTime, err := time.Parse(time.RFC3339, aTime)
		if err != nil { return nil, err }

		return models.Event{
			time.Now(),
			creatorId,
			place,
			activityId,
			parsedTime,
			lang}, nil
	},
}

func getErr(argName string, typeName string) error {
	return fmt.Errorf("%s can't be parsed into a %s", argName, typeName)
}