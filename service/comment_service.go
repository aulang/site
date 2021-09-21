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
	GetTop(num int64) ([]Comment, error)
	FindByArticleId(articleId string) ([]Comment, error)

	Reply(commentId string, reply *Reply) (Comment, error)
}

type commentService struct {
	c   *mongo.Collection
	ctx context.Context
}

func (s *commentService) GetTop(num int64) ([]Comment, error) {
	ops := options.Find().SetSort(bson.D{{"creationDate", -1}}).SetLimit(num)

	cur, err := s.c.Find(s.ctx, bson.D{}, ops)

	if err != nil {
		return nil, err
	}

	defer closeCursor(cur, s.ctx)

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
	ops := options.Find().SetSort(bson.D{{"creationDate", 1}})

	cur, err := s.c.Find(s.ctx, bson.D{{"articleId", articleId}}, ops)

	if err != nil {
		return nil, err
	}

	defer closeCursor(cur, s.ctx)

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

		_id, err := primitive.ObjectIDFromHex(comment.ArticleID)
		if err == nil {
			query := bson.D{{"_id", _id}}

			update := bson.D{
				{"$inc", bson.M{"commentsCount": 1}},
			}
			_, err = articleCollection.UpdateOne(s.ctx, query, update)

			return err
		}
		return err
	} else {
		_id := comment.ID

		query := bson.D{{"_id", _id}}

		update := bson.D{
			{"$set", comment},
		}

		_, err := s.c.UpdateOne(s.ctx, query, update)

		if err == mongo.ErrNoDocuments {
			return ErrNotFound
		}

		return err
	}
}

func (s *commentService) GetByID(id string) (Comment, error) {
	var comment Comment

	_id, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return comment, err
	}

	query := bson.D{{"_id", _id}}

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

	_, _ = commentCollection.Indexes().CreateMany(ctx, indexes[:])
}
