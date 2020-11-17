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

type CategoryService interface {
	GetAll() ([]Category, error)

	Save(category *Category) error
}

type categoryService struct {
	c   *mongo.Collection
	ctx context.Context
}

func (s *categoryService) getMaxOrder() int {
	var order = 1

	ops := options.FindOne().SetSort(bson.D{{Key: "order", Value: -1}})
	result := s.c.FindOne(s.ctx, bson.D{}, ops)

	if result.Err() != nil {
		return order
	}

	var elem Category
	err := result.Decode(&elem)

	if err != nil {
		return order
	}

	return elem.Order
}

func (s *categoryService) GetAll() ([]Category, error) {
	ops := options.Find().SetSort(bson.D{{Key: "order", Value: 1}})

	cur, err := s.c.Find(s.ctx, bson.D{}, ops)

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

func (s *categoryService) Save(category *Category) error {
	if category.ID.IsZero() {
		category.ID = primitive.NewObjectID()

		category.Order = s.getMaxOrder()

		_, err := s.c.InsertOne(s.ctx, category)

		if err != nil {
			return err
		}
	} else {
		_id := category.ID

		query := bson.D{{Key: "_id", Value: _id}}

		update := bson.D{
			{Key: "$set", Value: category},
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

var _ CategoryService = (*categoryService)(nil)

var categoryCollection = repository.Collection("category")

func NewCategoryService() CategoryService {
	return &categoryService{c: categoryCollection, ctx: ctx}
}

func init() {
	indexes := [...]mongo.IndexModel{
		{
			Keys:    bson.M{"name": 1},
			Options: options.Index().SetName("ik_category_name").SetUnique(true),
		},
	}

	categoryCollection.Indexes().CreateMany(ctx, indexes[:])
}
