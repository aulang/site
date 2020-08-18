package service

import (
	"context"

	. "github.com/aulang/site/entity"
	"github.com/aulang/site/model"
	"github.com/aulang/site/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ArticleService interface {
	GetAll() ([]Article, error)
	GetByID(id string) (Article, error)
	Save(a *Article) error
	Delete(id string) error
	GetTop3() ([]Article, error)
	Page(pageNo, pageSize int64, keyword, category string) (model.Page, error)
}

type articleService struct {
	c   *mongo.Collection
	ctx context.Context
}

func (s *articleService) GetAll() ([]Article, error) {
	cur, err := s.c.Find(s.ctx, bson.D{})

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

func (s *articleService) GetByID(id string) (Article, error) {
	var article Article

	_id, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return article, err
	}

	query := bson.D{{Key: "_id", Value: _id}}

	err = s.c.FindOne(s.ctx, query).Decode(&article)

	if err == mongo.ErrNoDocuments {
		return article, ErrNotFound
	}

	return article, err
}

func (s *articleService) Save(article *Article) error {
	if article.ID.IsZero() {
		article.ID = primitive.NewObjectID()
		_, err := s.c.InsertOne(s.ctx, category)

		if err != nil {
			return err
		}
	} else {
		_id := article.ID

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

func (s *articleService) Delete(id string) error {
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

func (s *articleService) GetTop3() ([]Article, error) {
	// 只要id和title字段
	ops := options.Find().SetProjection(bson.M{"_id": 1, "title": 1}).SetSort(bson.D{{Key: "creationDate", Value: -1}}).SetLimit(3)

	cur, err := s.c.Find(s.ctx, bson.D{}, ops)

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

func (s *articleService) Page(pageNo, pageSize int64, keyword, category string) (model.Page, error) {
	var page model.Page

	skip := (pageNo - 1) * pageSize

	ops := options.Find().SetSort(bson.D{{Key: "renew", Value: -1}}).SetSkip(skip).SetLimit(pageSize)

	filter := bson.M{}

	if category != "" {
		filter["categoryId"] = category
	}

	if keyword != "" {
		filter["$or"] = bson.A{
			bson.M{"title": bson.M{"$regex": keyword}},
			bson.M{"subTitle": bson.M{"$regex": keyword}},
		}
	}

	count, err := s.c.CountDocuments(s.ctx, filter)
	if err != nil {
		return page, err
	}

	cur, err := s.c.Find(s.ctx, filter, ops)

	if err != nil {
		return page, err
	}

	defer cur.Close(s.ctx)

	var datas []interface{}

	for cur.Next(s.ctx) {
		if err = cur.Err(); err != nil {
			return page, err
		}

		var elem Article
		err = cur.Decode(&elem)

		if err != nil {
			return page, err
		}

		datas = append(datas, elem)
	}

	page.Datas = datas
	page.PageNo = pageNo
	page.PageSize = pageNo
	page.TotalPages = (count + pageSize - 1) / pageSize

	return page, nil
}

var _ ArticleService = (*articleService)(nil)

var article = repository.Collection("article")

func NewArticleService() ArticleService {
	return &articleService{c: article, ctx: ctx}
}

func init() {
	indexes := [...]mongo.IndexModel{
		{
			Keys:    bson.M{"renew": -1},
			Options: options.Index().SetName("ik_article_renew").SetBackground(true),
		},
		{
			Keys:    bson.M{"creationDate": -1},
			Options: options.Index().SetName("ik_article_creationDate").SetBackground(true),
		},
		{
			Keys:    bson.M{"categoryId": -1},
			Options: options.Index().SetName("ik_article_categoryId").SetBackground(true),
		},
	}

	article.Indexes().CreateMany(ctx, indexes[:])
}
