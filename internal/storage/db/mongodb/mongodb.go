package mongodbStorage

import (
	"context"
	"errors"
	"fmt"

	"github.com/kuzja086/smartHome/internal/apperror"
	"github.com/kuzja086/smartHome/internal/entity"
	"github.com/kuzja086/smartHome/pkg/logging"

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
	d.logger.Info(err)
	return "", err
}

func (d *UserStorage) FindByUsername(ctx context.Context, username string) (u entity.User, err error) {
	filter := bson.M{"username": username}

	result := d.collection.FindOne(ctx, filter)
	if result.Err() != nil {
		d.logger.Error(result.Err())
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			return u, apperror.UserNotFound
		}
		return u, fmt.Errorf("failed to execute query. error: %w", err)
	}

	if err = result.Decode(&u); err != nil {
		return u, err
	}

	return u, nil
}
