package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	. "site/entity"
)

type BeiAnService interface {
	Get() (BeiAn, error)
}

type beiAnService struct {
	C   *mongo.Collection
	ctx context.Context
}

func (s *beiAnService) Get() (BeiAn, error) {
	var beiAn BeiAn

	err := s.C.FindOne(s.ctx, bson.D{}).Decode(&beiAn)

	if err == mongo.ErrNoDocuments {
		return beiAn, ErrNotFound
	}

	return beiAn, nil
}

var _ BeiAnService = (*beiAnService)(nil)

func NewBeiAnService(collection *mongo.Collection) BeiAnService {
	return &beiAnService{C: collection, ctx: context.Background()}
}
