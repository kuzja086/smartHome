package mongodbStorage

import (
	"context"
	"smartHome/internal/entity"
	"smartHome/pkg/logging"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserStorage struct {
	collection *mongo.Collection
	logger     *logging.Logger
}

func NewUserStorage(database *mongo.Database, collection string, logger *logging.Logger) *UserStorage {
	return &UserStorage{
		collection: database.Collection(collection),
		logger:     logger,
	}
}

func (d *UserStorage) CreateUser(ctx context.Context, user entity.User) (string, error) {
	d.logger.Debug("create user")
	result, err := d.collection.InsertOne(ctx, user)
	if err != nil {
		return "", err
	}

	d.logger.Debug("convert insertetID in ObjectID")
	oid, ok := result.InsertedID.(primitive.ObjectID)
	if ok {
		return oid.Hex(), nil
	}
	d.logger.Trace(user)
	return "", err
}

func (d *UserStorage) FindOne(ctx context.Context, id string) (u entity.User, err error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		d.logger.Debugf("HEX: %s", id)
		return u, err
	}

	filter := bson.M{"_id": oid}

	result := d.collection.FindOne(ctx, filter)
	if result.Err() != nil {
		// TODO 404
		return u, result.Err()
	}

	if err = result.Decode(&u); err != nil {
		return u, err
	}

	return u, nil
}
