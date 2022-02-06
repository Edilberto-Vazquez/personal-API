package services

import (
	"context"

	"github.com/Edilberto-Vazquez/personal-API/libs"
	"go.mongodb.org/mongo-driver/bson"
)

func AboutMeFindOne(alias string) (result bson.M, err error) {
	coll := libs.Client.Database("mydata").Collection("about_me")
	err = coll.FindOne(context.TODO(), bson.D{{"alias", alias}}).Decode(&result)
	return result, err
}
