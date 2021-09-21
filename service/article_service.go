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
	GetByID(id string) (Article, error)
	Save(article *Article) error
	Delete(id string) error
	GetTop(num int64) ([]Article, error)
	GetByPage(page, size int64, keyword, category string) (*model.Page, error)
}

type articleService struct {
	c   *mongo.Collection
	ctx context.Context
}

func (s *articleService) GetByID(id string) (Article, error) {
	var article Article

	_id, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return article, err
	}

	query := bson.D{{"_id", _id}}

	err = s.c.FindOne(s.ctx, query).Decode(&article)

	if err == mongo.ErrNoDocuments {
		return article, ErrNotFound
	}

	return article, err
}

func (s *articleService) Save(article *Article) error {
	if article.ID.IsZero() {
		article.ID = primitive.NewObjectID()
		_, err := s.c.InsertOne(s.ctx, article)

		if err != nil {
			return err
		}

		_id, err := primitive.ObjectIDFromHex(article.CategoryID)
		if err == nil {
			query := bson.D{{"_id", _id}}

			update := bson.D{
				{Key: "$inc", Value: bson.M{"count": 1}},
			}
			_, err = categoryCollection.UpdateOne(s.ctx, query, update)

			return err
		}
		return err
	} else {
		_id := article.ID

		query := bson.D{{"_id", _id}}

		update := bson.D{
			{"$set", article},
		}

		_, err := s.c.UpdateOne(s.ctx, query, update)

		if err == mongo.ErrNoDocuments {
			return ErrNotFound
		}

		return err
	}
}

func (s *articleService) Delete(id string) error {
	_id, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	query := bson.D{{"_id", _id}}

	_, err = s.c.DeleteOne(s.ctx, query)

	if err == mongo.ErrNoDocuments {
		return ErrNotFound
	}

	return err
}

func (s *articleService) GetTop(num int64) ([]Article, error) {
	// 只要id和title字段
	ops := options.Find().SetProjection(bson.M{"_id": 1, "title": 1}).SetSort(bson.D{{"creationDate", -1}}).SetLimit(num)

	cur, err := s.c.Find(s.ctx, bson.D{}, ops)

	if err != nil {
		return nil, err
	}

	defer closeCursor(cur, s.ctx)

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

func (s *articleService) GetByPage(pageNo, pageSize int64, keyword, category string) (*model.Page, error) {
	skip := (pageNo - 1) * pageSize

	ops := options.Find().SetSort(bson.D{{"renew", -1}}).SetSkip(skip).SetLimit(pageSize)

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
		return nil, err
	}

	cur, err := s.c.Find(s.ctx, filter, ops)

	if err != nil {
		return nil, err
	}

	defer closeCursor(cur, s.ctx)

	var articles []interface{}

	for cur.Next(s.ctx) {
		if err = cur.Err(); err != nil {
			return nil, err
		}

		var elem Article
		err = cur.Decode(&elem)

		if err != nil {
			return nil, err
		}

		articles = append(articles, elem)
	}

	page := model.NewPage(pageNo, pageSize, count, articles)

	return page, nil
}

var _ ArticleService = (*articleService)(nil)

var articleCollection = repository.Collection("article")

func NewArticleService() ArticleService {
	return &articleService{c: articleCollection, ctx: ctx}
}

func init() {
	indexes := [...]mongo.IndexModel{
		{
			Keys:    bson.M{"renew": -1},
			Options: options.Index().SetName("ik_article_renew"),
		},
		{
			Keys:    bson.M{"creationDate": -1},
			Options: options.Index().SetName("ik_article_creationDate"),
		},
		{
			Keys:    bson.M{"categoryId": -1},
			Options: options.Index().SetName("ik_article_categoryId"),
		},
	}

	_, _ = articleCollection.Indexes().CreateMany(ctx, indexes[:])
}
