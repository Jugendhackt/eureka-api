package queries

import (
	"github.com/graphql-go/graphql"
	"../../models"
	gqlModels "../models"
	"../../helpers"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"fmt"
)

var User = graphql.Field{
	Type: gqlModels.User,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		user := models.User{}
		var err error = nil

		helpers.SetUpDb(func(db *mgo.Database) {
			id, isOk := params.Args["id"].(string)

			if !isOk {
				err = fmt.Errorf("could not parse id to string")
				return
			}

			fmt.Println(id)

			err = db.C("users").FindId(bson.ObjectIdHex(id)).One(&user)
		})
		return user, err
	},
}