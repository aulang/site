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

type MenuService interface {
	GetAll() ([]Menu, error)

	Save(menu *Menu) error
}

type menuService struct {
	c   *mongo.Collection
	ctx context.Context
}

func (s *menuService) getMaxOrder() int {
	var order int = 1

	ops := options.FindOne().SetSort(bson.D{{Key: "order", Value: -1}})
	result := s.c.FindOne(s.ctx, bson.D{}, ops)

	if result.Err() != nil {
		return order
	}

	var elem Menu
	err := result.Decode(&elem)

	if err != nil {
		return order
	}

	return elem.Order
}

func (s *menuService) GetAll() ([]Menu, error) {
	ops := options.Find().SetSort(bson.D{{Key: "order", Value: 1}})

	cur, err := s.c.Find(s.ctx, bson.D{}, ops)

	if err != nil {
		return nil, err
	}

	defer cur.Close(s.ctx)

	var results []Menu

	for cur.Next(s.ctx) {
		if err = cur.Err(); err != nil {
			return nil, err
		}

		var elem Menu
		err = cur.Decode(&elem)

		if err != nil {
			return nil, err
		}

		results = append(results, elem)
	}

	return results, nil
}

func (s *menuService) Save(menu *Menu) error {
	if menu.ID.IsZero() {
		menu.ID = primitive.NewObjectID()

		menu.Order = s.getMaxOrder()

		_, err := s.c.InsertOne(s.ctx, menu)

		if err != nil {
			return err
		}
	} else {
		_id := menu.ID

		query := bson.D{{Key: "_id", Value: _id}}

		update := bson.D{
			{Key: "$set", Value: menu},
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

var _ MenuService = (*menuService)(nil)

var menu = repository.Collection("menu")

func NewMenuService() MenuService {
	return &menuService{c: menu, ctx: ctx}
}
