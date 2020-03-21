package scorer

import (
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
	"log"
	"os"
)

var client *db.Client

func init() {
	ctx := context.Background()
	dbURL := os.Getenv("firebaseDBURL")
	conf := &firebase.Config{
		DatabaseURL: dbURL,
	}

	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		log.Panic("error creating new app using firebase, err:" + err.Error())
	}

	client, err = app.Database(ctx)
	if err != nil {
		log.Panic("error creating database client, err:" + err.Error())
	}
}
