package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	. "site/entity"
	"site/repository"
)

type MenuService interface {
	GetAll() ([]Menu, error)

	Save(a *Menu) error
}

type menuService struct {
	c   *mongo.Collection
	ctx context.Context
}

func (s *menuService) GetAll() ([]Menu, error) {
	cur, err := s.c.Find(s.ctx, bson.D{})

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

func (s *menuService) Save(m *Menu) error {
	if m.ID.IsZero() {
		m.ID = primitive.NewObjectID()
		_, err := s.c.InsertOne(s.ctx, m)

		if err != nil {
			return err
		}
	} else {
		_id := m.ID

		query := bson.D{{Key: "_id", Value: _id}}

		update := bson.D{
			{Key: "$set", Value: m},
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
