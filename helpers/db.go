package helpers

import (
	"gopkg.in/mgo.v2"
	"os"
	"log"
)

func SetUpDb(body func(db *mgo.Database)) {
	session, err := mgo.Dial(os.Getenv("DB_URL"))
	defer session.Close()

	if err != nil { log.Fatal(err)}

	session.SetMode(mgo.Monotonic, true)
	db := session.DB(os.Getenv("DB_NAME"))
	body(db)
}