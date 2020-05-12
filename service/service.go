package service

import (
	"context"
	"errors"
)

var ctx = context.Background()

var ErrNotFound = errors.New("记录不存在")
