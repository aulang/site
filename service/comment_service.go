package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	. "site/entity"
	"site/repository"
)

type CommentService interface {
	GetTop3() ([]Comment, error)
}

type commentService struct {
	C   *mongo.Collection
	ctx context.Context
}

func (s *commentService) GetTop3() ([]Comment, error) {
	// 只要id和title字段
	ops := options.Find().SetSort(bson.D{{Key: "creationDate", Value: -1}}).SetLimit(3)

	cur, err := s.C.Find(s.ctx, bson.D{}, ops)

	if err != nil {
		return nil, err
	}

	defer cur.Close(s.ctx)

	var results []Comment

	for cur.Next(s.ctx) {
		if err = cur.Err(); err != nil {
			return nil, err
		}

		var elem Comment
		err = cur.Decode(&elem)

		if err != nil {
			return nil, err
		}

		results = append(results, elem)
	}

	return results, nil
}

var _ CommentService = (*commentService)(nil)

func NewCommentService() CommentService {
	collection := repository.Collection("comment")
	return &commentService{C: collection, ctx: context.Background()}
}