package services

import (
	"context"

	"github.com/Edilberto-Vazquez/personal-API/libs"
	"github.com/Edilberto-Vazquez/personal-API/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func ExperienceAndEducationFind(resumeType string) (results []bson.M, err error) {
	var cursor *mongo.Cursor
	coll := libs.Client.Database("my-api").Collection(resumeType)
	cursor, err = coll.Find(context.TODO(), bson.D{{}})
	err = utils.CursorDecode(&results, cursor, err)
	return results, err
}
