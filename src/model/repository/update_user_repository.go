package repository

import (
	"context"
	"os"

	rest_err "github.com/DavibernardesA/CRUD-Go/src/configuration/api_err"
	"github.com/DavibernardesA/CRUD-Go/src/configuration/logger"
	"github.com/DavibernardesA/CRUD-Go/src/model"
	"github.com/DavibernardesA/CRUD-Go/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (ur *userRepository) UpdateUser(
	userId string,
	userDomain model.UserDomainInterface,
) *rest_err.ApiErr {
	logger.Info("Init updateUser repository",
		zap.String("journey", "updateUser"))

	collection_name := os.Getenv(MONGODB_USER)
	collection := ur.databaseConnection.Collection(collection_name)

	value := converter.ConvertDomainToEntity(userDomain)
	userIdHex, _ := primitive.ObjectIDFromHex(userId)

	filter := bson.D{{Key: "_id", Value: userIdHex}}
	update := bson.D{{Key: "$set", Value: value}}

	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		logger.Error("Error trying to update user",
			err,
			zap.String("journey", "updateUser"))
		return rest_err.NewInternalServerError(err.Error())
	}

	logger.Info(
		"updateUser repository executed successfully",
		zap.String("userId", userId),
		zap.String("journey", "updateUser"))

	return nil
}
