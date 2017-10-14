package queries

import (
	"github.com/graphql-go/graphql"
	gqlModels "../models"
	"../../models"
	"../../helpers"
	"gopkg.in/mgo.v2"
	"fmt"
	"gopkg.in/mgo.v2/bson"
)

var Event = graphql.Field{
	Type: gqlModels.Event,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		var event models.Event
		var err error = nil

		id, isOk := params.Args["id"].(string)
		if !isOk { return event, fmt.Errorf("could not parse id to string")}

		helpers.SetUpDb(func(db *mgo.Database) {
			err = db.C("events").FindId(bson.ObjectIdHex(id)).One(&event)
		})

		return event, err
	},
}