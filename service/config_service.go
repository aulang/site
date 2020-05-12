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

var _ WebConfigService = (*webConfigService)(nil)

var webConfig = repository.Collection("webConfig")

func NewWebConfigService() WebConfigService {
	return &webConfigService{c: webConfig, ctx: ctx}
}
