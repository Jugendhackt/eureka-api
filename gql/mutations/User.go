package mutations

import (
	"github.com/graphql-go/graphql"
	gqlModels "../models"
	"../../models"
	"fmt"
	"time"
	"gopkg.in/mgo.v2"
	"log"
	"../../helpers"
	"gopkg.in/mgo.v2/bson"
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
		helpers.SetUpDb(func(db *mgo.Database) {
			err := db.C("users").Insert(&user)
			if err != nil { log.Fatalf("%v", err) }
		})

		return user, nil
	},
}

var UpdateUser = graphql.Field{
	Type: gqlModels.User,
	Args: graphql.FieldConfigArgument{
		// REQUIRED: to know which user
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		// OPTIONALS: Don't have to be changed
		"emailAddress": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"hometown": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		// TODO: FIgure out how to EXACTLY up (Without knowing value; it's possible)
		//"experiencePoints": &graphql.ArgumentConfig{ // Will up/lower current state by value
		//	Type: graphql.Int,
		//},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {

		user := models.User{}
		userUpdateMap := make(map[string]interface{})

		idStr, isOk := params.Args["id"].(string)
		if !isOk { return nil, fmt.Errorf("couldn't parse id (to string)") }

		emailAddress, isOk := params.Args["emailAddress"].(string)
		if isOk {
			userUpdateMap["emailAddress"] = emailAddress
		}

		name, isOk := params.Args["name"].(string)
		if isOk {
			userUpdateMap["name"] = name
		}

		hometown, isOk := params.Args["hometown"].(string)
		if isOk {
			userUpdateMap["hometown"] = hometown
		}

		helpers.SetUpDb(func(db *mgo.Database) {
			id := bson.ObjectIdHex(idStr)
			db.C("users").UpdateId(id, bson.M{"$set": userUpdateMap})
		})
		// TODO: What should I return here?
		return user, nil
	},
}