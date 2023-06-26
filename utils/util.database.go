package utils

import (
	"context"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

var client *firestore.Client

func GetFirstoreClient() *firestore.Client {
	return client
}

func CloseFirestoreClient() error {
	return client.Close()
}

func InitiateFirestoreClient() error {
	ctx := context.Background()

	option := option.WithCredentialsFile("./config/firestore.json")
	app, err := firebase.NewApp(ctx, nil, option)
	if err != nil {
		return err
	}

	client, err = app.Firestore(ctx)
	if err != nil {
		return err
	}

	return nil
}

func GetDocumentRefByPath(path string) *firestore.DocumentRef {
	return client.Doc(path)
}