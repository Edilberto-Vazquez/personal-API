package utils

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CursorDecode(results *[]bson.M, cursor *mongo.Cursor, err error) error {
	if err != nil {
		return err
	}
	if err = cursor.All(context.TODO(), results); err != nil {
		cursor.Close(context.TODO())
		return err
	}
	return err
}
