package queries

import (
	"github.com/graphql-go/graphql"
	gqlModels "../models"
	"../../models"
	"../../helpers"
	"gopkg.in/mgo.v2"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"time"
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

var Events = graphql.Field{
	Type: gqlModels.EventResults,
	Args: graphql.FieldConfigArgument{
		"offset": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		var err error = nil

		offset, isOk := params.Args["offset"].(int)
		if !isOk {
			offset = 0
		}
		var recent []models.Event
		var upcoming []models.Event

		helpers.SetUpDb(func(db *mgo.Database) {
			// Get most recent
			err = db.C("events").Find(nil).Sort("-creationDate").Skip(offset).Limit(10).All(&recent)
			if err != nil {
				return
			}
			err = db.C("events").Find(bson.M{"time": bson.M{"$lt": time.Now()}}).Limit(10).All(&upcoming)
		})

		fmt.Println(upcoming)

		result := models.EventResults{
			recent,
			upcoming,
		}

		return result, err
	},
}