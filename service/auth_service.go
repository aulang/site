package service

import (
	"context"
	"github.com/aulang/site/entity"
	"github.com/aulang/site/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type AuthService interface {
	Get() entity.Auth
}

type authService struct {
	c   *mongo.Collection
	ctx context.Context
}

func (s *authService) Get() entity.Auth {
	var auth entity.Auth

	err := s.c.FindOne(ctx, bson.D{}).Decode(&auth)
	if err != nil {
		log.Fatalln("未配置认证信息，系统退出！")
	}

	return auth
}

var auth = repository.Collection("auth")

func NewAuthService() AuthService {
	return &authService{c: auth, ctx: ctx}
}
