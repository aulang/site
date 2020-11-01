package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"

	. "github.com/aulang/site/entity"
	"github.com/aulang/site/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type WebConfigService interface {
	Get() (WebConfig, error)
	Save(config *WebConfig) error
}

type webConfigService struct {
	c   *mongo.Collection
	ctx context.Context
}

func (s *webConfigService) Get() (WebConfig, error) {
	var webConfig WebConfig

	err := s.c.FindOne(s.ctx, bson.D{}).Decode(&webConfig)

	if err == mongo.ErrNoDocuments {
		return webConfig, ErrNotFound
	}

	return webConfig, nil
}

func (s *webConfigService) Save(config *WebConfig) error {
	if config.ID.IsZero() {
		config.ID = primitive.NewObjectID()

		_, err := s.c.InsertOne(s.ctx, config)

		if err != nil {
			return err
		}
	} else {
		_id := config.ID

		query := bson.D{{Key: "_id", Value: _id}}

		update := bson.D{
			{Key: "$set", Value: config},
		}

		_, err := s.c.UpdateOne(s.ctx, query, update)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				return ErrNotFound
			}

			return err
		}
	}

	return nil
}

var _ WebConfigService = (*webConfigService)(nil)

var webConfigCollection = repository.Collection("webConfig")

func NewWebConfigService() WebConfigService {
	return &webConfigService{c: webConfigCollection, ctx: ctx}
}
