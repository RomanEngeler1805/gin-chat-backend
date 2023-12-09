package services

import (
    "context"
    "cloud.google.com/go/firestore"
    "google.golang.org/api/option"
)

func GetChat(id string) (interface{}, error) {
    ctx := context.Background()
	// TODO: replace this with the location of the service account
    sa := option.WithCredentialsFile("/Users/romanengeler/gin-chat-backend/firestore_key.json")
    client, err := firestore.NewClient(ctx, "django-chat-app-firestore", sa)
    if err != nil {
        return nil, err
    }
    defer client.Close()

    doc, err := client.Collection("chat").Doc(id).Get(ctx)
    if err != nil {
        return nil, err
    }
    return doc.Data(), nil
}