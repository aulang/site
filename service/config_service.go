package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	. "site/entity"
	"site/repository"
)

type WebConfigService interface {
	Get() (WebConfig, error)
}

type webConfigService struct {
	C   *mongo.Collection
	ctx context.Context
}

func (s *webConfigService) Get() (WebConfig, error) {
	var webConfig WebConfig

	err := s.C.FindOne(s.ctx, bson.D{}).Decode(&webConfig)

	if err == mongo.ErrNoDocuments {
		return webConfig, ErrNotFound
	}

	return webConfig, nil
}

var _ WebConfigService = (*webConfigService)(nil)

func NewWebConfigService() WebConfigService {
	collection := repository.Collection("webConfig")
	return &webConfigService{C: collection, ctx: context.Background()}
}
