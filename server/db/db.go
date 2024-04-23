package db

import (
	"context"
	"log"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)


const (
    firebaseConfigFile = "../to-do-list-f2708-firebase-adminsdk-4yqgh-2974d9037b.json"
    firebaseDBURL      = "https://To-Do-List.firebaseio.com/"
)

var (
    ctx context.Context
    app *firebase.App
)

func init() {
	ctx := context.Background()
	opt := option.WithCredentialsFile(firebaseConfigFile)
    app, err := firebase.NewApp(ctx, nil, opt)
    if err != nil {
        log.Fatalf("Firebase initialization error: %v\n", err)
    }

	client, err := app.DatabaseWithURL(ctx, firebaseDBURL)
	
    if err != nil {
        log.Fatalf("Firestore initialization error: %v\n", err)
    }
}