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
	FindByArticleId(articleId string) ([]Comment, error)
}

type commentService struct {
	c   *mongo.Collection
	ctx context.Context
}

func (s *commentService) GetTop3() ([]Comment, error) {
	ops := options.Find().SetSort(bson.D{{Key: "creationDate", Value: -1}}).SetLimit(3)

	cur, err := s.c.Find(s.ctx, bson.D{}, ops)

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

func (s *commentService) FindByArticleId(articleId string) ([]Comment, error) {
	ops := options.Find().SetSort(bson.D{{Key: "creationDate", Value: 1}})

	cur, err := s.c.Find(s.ctx, bson.D{{"articleId", articleId}}, ops)

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

var comment = repository.Collection("comment")

func NewCommentService() CommentService {
	return &commentService{c: comment, ctx: ctx}
}

func init() {
	indexes := [...]mongo.IndexModel{
		{
			Keys:    bson.M{"comment": -1},
			Options: options.Index().SetName("ik_comment_articleId").SetBackground(true),
		},
	}

	comment.Indexes().CreateMany(ctx, indexes[:])
}
