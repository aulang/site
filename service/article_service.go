package service

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	. "site/entity"
)

type ArticleService interface {
	GetAll() ([]Article, error)
	GetByID(id string) (Article, error)
	Create(a *Article) error
	Update(a Article) error
	Delete(id string) error
}

type articleService struct {
	C   *mongo.Collection
	ctx context.Context
}

func (s *articleService) GetAll() ([]Article, error) {
	cur, err := s.C.Find(s.ctx, bson.D{})

	if err != nil {
		return nil, err
	}

	defer cur.Close(s.ctx)

	var results []Article

	for cur.Next(s.ctx) {
		if err = cur.Err(); err != nil {
			return nil, err
		}

		var elem Article
		err = cur.Decode(&elem)

		if err != nil {
			return nil, err
		}

		results = append(results, elem)
	}

	return results, nil
}

var ErrNotFound = errors.New("记录不存在")

func (s *articleService) GetByID(id string) (Article, error) {
	var article Article

	_id, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return article, err
	}

	query := bson.D{{Key: "_id", Value: _id}}

	err = s.C.FindOne(s.ctx, query).Decode(&article)

	if err == mongo.ErrNoDocuments {
		return article, ErrNotFound
	}

	return article, err
}

func (s *articleService) Create(a *Article) error {
	if a.ID.IsZero() {
		a.ID = primitive.NewObjectID()
	}

	_, err := s.C.InsertOne(s.ctx, a)

	if err != nil {
		return err
	}

	return nil
}

func (s *articleService) Update(a Article) error {
	_id := a.ID

	query := bson.D{{Key: "_id", Value: _id}}

	update := bson.D{
		{Key: "$set", Value: a},
	}

	_, err := s.C.UpdateOne(s.ctx, query, update)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return ErrNotFound
		}

		return err
	}

	return nil
}

func (s *articleService) Delete(id string) error {
	_id, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	query := bson.D{{Key: "_id", Value: _id}}

	_, err = s.C.DeleteOne(s.ctx, query)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return ErrNotFound
		}

		return err
	}

	return nil
}

var _ ArticleService = (*articleService)(nil)

func NewArticleService(collection *mongo.Collection) ArticleService {
	return &articleService{C: collection, ctx: context.Background()}
}
