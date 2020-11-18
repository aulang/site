package service

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

var ctx = context.Background()

var ErrNotFound = errors.New("记录不存在")

func closeCursor(cur *mongo.Cursor, ctx context.Context) {
	err := cur.Close(ctx)
	if err != nil {
		log.Println("关闭游标失败：", err)
	}
}
