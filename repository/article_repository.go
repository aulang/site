package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"site/entity"
)

type ArticleRepository struct {
	collectionName string
}

func NewArticleRepository() *ArticleRepository {
	return &ArticleRepository{
		collectionName: "article",
	}
}

func (s *ArticleRepository) CollectionName() string {
	return s.collectionName
}

func (s *ArticleRepository) SelectMany(query bson.D, skip int64, limit int64) (results []entity.Article) {
	cur, err := Collection(s.collectionName).Find(context.TODO(), query, &options.FindOptions{
		Limit: &limit,
		Skip:  &skip,
		Sort:  bson.M{"renew": -1},
	})

	if err != nil {
		panic(err)
	}

	for cur.Next(context.TODO()) {
		var elem entity.Article
		err := cur.Decode(&elem)
		if err != nil {
			panic(err)
		}

		results = append(results, elem)
	}

	if err := cur.Err(); err != nil {
		panic(err)
	}

	cur.Close(context.TODO())

	return results
}
