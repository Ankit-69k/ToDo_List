package db

import (
	"context"
	"encoding/json"
	"os"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

var firestoreClient *firestore.Client

func getFirebaseCredentials() ([]byte, error) {
	err := godotenv.Load() // Load .env file
	if err != nil {
		return nil, err // Return error, do not terminate the application
	}

	serviceAccount := map[string]interface{}{
		"type":                        os.Getenv("FIREBASE_TYPE"),
		"project_id":                  os.Getenv("FIREBASE_PROJECT_ID"),
		"private_key_id":              os.Getenv("FIREBASE_PRIVATE_KEY_ID"),
		"private_key":                 os.Getenv("FIREBASE_PRIVATE_KEY"),
		"client_email":                os.Getenv("FIREBASE_CLIENT_EMAIL"),
		"client_id":                   os.Getenv("FIREBASE_CLIENT_ID"),
		"auth_uri":                    os.Getenv("FIREBASE_AUTH_URI"),
		"token_uri":                   os.Getenv("FIREBASE_TOKEN_URI"),
		"auth_provider_x509_cert_url": os.Getenv("FIREBASE_AUTH_PROVIDER_X509_CERT_URL"),
		"client_x509_cert_url":        os.Getenv("FIREBASE_CLIENT_X509_CERT_URL"),
	}

	credentials, err := json.Marshal(serviceAccount)
	if err != nil {
		return nil, err // Return the error, no termination
	}

	return credentials, nil
}

func InitFirestore() (*Store, error)  {
	
	ctx := context.Background()
	creds, err := getFirebaseCredentials() // Get credentials
	
	if err != nil {
		return nil, err // Return error if credentials cannot be retrieved
	}

	opt := option.WithCredentialsJSON(creds)
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, err
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		return nil, err
	}

	 // Store the Firestore client in a global variable
	return &Store{client:client}, nil
}
