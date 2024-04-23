package auth

import (
	"context"
	"fmt"
	"log"
	"net/http"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
	"firebase.google.com/go/v4/auth"
	"cloud.google.com/go/firestore"
)

var app *firebase.App
var authClient *auth.Client
var firestoreClient *firestore.Client

func initFirebase() {
	ctx := context.Background()

	credentialsFilePath := "../to-do-list-f2708-firebase-adminsdk-4yqgh-2974d9037b.json"
	opt := option.WithCredentialsFile(credentialsFilePath)

	var err error
	app, err = firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("Error initializing Firebase app: %v", err)
	}

	authClient, err = app.Auth(ctx)
	if err != nil {
		log.Fatalf("Error initializing Firebase Auth client: %v", err)
	}

	firestoreClient, err = app.Firestore(ctx)
	if err != nil {
		log.Fatalf("Error initializing Firestore client: %v", err)
	}

	fmt.Println("Firebase initialized successfully!")
}

func validateToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header is missing", http.StatusUnauthorized)
			return
		}

		tokenStr := ""
		if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
			tokenStr = authHeader[7:]
		} else {
			http.Error(w, "Invalid authorization header format", http.StatusUnauthorized)
			return
		}

		token, err := authClient.VerifyIDToken(r.Context(), tokenStr)
		if err != nil {
			http.Error(w, "Invalid ID token", http.StatusUnauthorized)
			return
		}

		// Add the user UID to the request context for later use
		ctx := context.WithValue(r.Context(), "userUID", token.UID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}