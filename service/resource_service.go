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

type ResourceService interface {
	GetByID(id string) (Resource, error)
	GetBySubjectID(subjectId string) ([]Resource, error)
	Save(resource *Resource) error
	Delete(id string) error
}

type resourceService struct {
	c   *mongo.Collection
	ctx context.Context
}

func (s *resourceService) GetByID(id string) (Resource, error) {
	var resource Resource

	_id, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return resource, err
	}

	query := bson.D{{Key: "_id", Value: _id}}

	err = s.c.FindOne(s.ctx, query).Decode(&resource)

	if err == mongo.ErrNoDocuments {
		return resource, ErrNotFound
	}

	return resource, err
}

func (s *resourceService) GetBySubjectID(subjectId string) ([]Resource, error) {
	query := bson.D{{Key: "subjectId", Value: subjectId}}

	cur, err := s.c.Find(s.ctx, query)

	if err != nil {
		return nil, err
	}

	defer closeCursor(cur, s.ctx)

	var results []Resource

	for cur.Next(s.ctx) {
		if err = cur.Err(); err != nil {
			return nil, err
		}

		var elem Resource
		err = cur.Decode(&elem)

		if err != nil {
			return nil, err
		}

		results = append(results, elem)
	}

	return results, nil
}

func (s *resourceService) Save(resource *Resource) error {
	if resource.ID.IsZero() {
		resource.ID = primitive.NewObjectID()

		_, err := s.c.InsertOne(s.ctx, resource)

		if err != nil {
			return err
		}
	} else {
		_id := resource.ID

		query := bson.D{{Key: "_id", Value: _id}}

		update := bson.D{
			{Key: "$set", Value: resource},
		}

		_, err := s.c.UpdateOne(s.ctx, query, update, options.Update().SetUpsert(true))
		if err != nil {
			if err == mongo.ErrNoDocuments {
				return ErrNotFound
			}

			return err
		}
	}

	return nil
}

func (s *resourceService) Delete(id string) error {
	_id, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	query := bson.D{{Key: "_id", Value: _id}}

	_, err = s.c.DeleteOne(s.ctx, query)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return ErrNotFound
		}

		return err
	}

	return nil
}

var _ ResourceService = (*resourceService)(nil)

var resourceCollection = repository.Collection("resource")

func NewResourceService() ResourceService {
	return &resourceService{c: resourceCollection, ctx: ctx}
}

// 创建索引
func init() {
	indexes := [...]mongo.IndexModel{
		{
			Keys:    bson.M{"subjectId": -1},
			Options: options.Index().SetName("ik_subjectId"),
		},
	}

	_, _ = resourceCollection.Indexes().CreateMany(ctx, indexes[:])
}
