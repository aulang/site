package service

import (
	"context"

	. "github.com/aulang/site/entity"
	"github.com/aulang/site/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CommentService interface {
	Save(comment *Comment) error
	GetTop3() ([]Comment, error)
	FindByArticleId(articleId string) ([]Comment, error)

	Reply(commentId string, reply *Reply) (Comment, error)
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

	cur, err := s.c.Find(s.ctx, bson.D{{Key: "articleId", Value: articleId}}, ops)

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

func (s *commentService) Save(comment *Comment) error {
	if comment.ID.IsZero() {
		comment.ID = primitive.NewObjectID()
		_, err := s.c.InsertOne(s.ctx, comment)

		if err != nil {
			return err
		}
	} else {
		_id := comment.ID

		query := bson.D{{Key: "_id", Value: _id}}

		update := bson.D{
			{Key: "$set", Value: comment},
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

func (s *commentService) GetByID(id string) (Comment, error) {
	var comment Comment

	_id, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return comment, err
	}

	query := bson.D{{Key: "_id", Value: _id}}

	err = s.c.FindOne(s.ctx, query).Decode(&comment)

	if err == mongo.ErrNoDocuments {
		return comment, ErrNotFound
	}

	return comment, err
}

func (s *commentService) Reply(commentId string, reply *Reply) (Comment, error) {
	comment, err := s.GetByID(commentId)
	if err != nil {
		return comment, err
	}

	replies := comment.Replies

	if replies == nil {
		replies = make([]Reply, 1, 1)
	}

	replies = append(replies, *reply)

	err = s.Save(&comment)

	return comment, err
}

var _ CommentService = (*commentService)(nil)

var commentCollection = repository.Collection("comment")

func NewCommentService() CommentService {
	return &commentService{c: commentCollection, ctx: ctx}
}

func init() {
	indexes := [...]mongo.IndexModel{
		{
			Keys:    bson.M{"articleId": -1},
			Options: options.Index().SetName("ik_comment_articleId"),
		},
	}

	commentCollection.Indexes().CreateMany(ctx, indexes[:])
}
