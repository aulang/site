package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	. "site/entity"
	"site/repository"
)

type CategoryService interface {
	GetAll() ([]Category, error)

	Save(a *Category) error
}

type categoryService struct {
	C   *mongo.Collection
	ctx context.Context
}

func (s *categoryService) GetAll() ([]Category, error) {
	cur, err := s.C.Find(s.ctx, bson.D{})

	if err != nil {
		return nil, err
	}

	defer cur.Close(s.ctx)

	var results []Category

	for cur.Next(s.ctx) {
		if err = cur.Err(); err != nil {
			return nil, err
		}

		var elem Category
		err = cur.Decode(&elem)

		if err != nil {
			return nil, err
		}

		results = append(results, elem)
	}

	return results, nil
}

func (s *categoryService) Save(m *Category) error {
	if m.ID.IsZero() {
		m.ID = primitive.NewObjectID()
		_, err := s.C.InsertOne(s.ctx, m)

		if err != nil {
			return err
		}
	} else {
		_id := m.ID

		query := bson.D{{Key: "_id", Value: _id}}

		update := bson.D{
			{Key: "$set", Value: m},
		}

		_, err := s.C.UpdateOne(s.ctx, query, update)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				return ErrNotFound
			}

			return err
		}
	}

	return nil
}

var _ CategoryService = (*categoryService)(nil)

func NewCategoryService() CategoryService {
	collection := repository.Collection("category")
	return &categoryService{C: collection, ctx: context.Background()}
}
