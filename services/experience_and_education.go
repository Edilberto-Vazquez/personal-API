package services

import (
	"context"

	"github.com/Edilberto-Vazquez/personal-API/libs"
	"go.mongodb.org/mongo-driver/bson"
)

func ExperienceAndEducationFind(resumeType string) (results []bson.M, err error) {
	coll := libs.Client.Database("mydata").Collection(resumeType)
	cursor, err := coll.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return results, err
	}
	if err = cursor.All(context.TODO(), &results); err != nil {
		cursor.Close(context.TODO())
		return results, err
	}
	return results, err
}
