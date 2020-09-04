package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Auth struct {
	ID           primitive.ObjectID `json:"id" bson:"_id"`
	ClientId     string             `json:"clientId" bson:"clientId"`
	ClientSecret string             `json:"clientSecret" bson:"clientId"`
	JwtSignKey   string             `json:"jwtSignKey" bson:"jwtSignKey"`
	JwtEncKey    string             `json:"jwtEncKey" bson:"jwtEncKey"`
}
