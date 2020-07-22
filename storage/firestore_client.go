package storage

import (
	"context"

	"cloud.google.com/go/firestore"
)

const (
	projectId string = "pragmatic-reviews-8060a"
)

func GetDBClient(ctx context.Context) (*firestore.Client, error) {
	return firestore.NewClient(ctx, projectId)
}
