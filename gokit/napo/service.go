package napo

import (
	"context"
	"time"
)

// 服务为您的应用程序提供了一些「日期功能」
type Service interface {
	Status(ctx context.Context) (string, error)
	Get(ctx context.Context) (string, error)
	Validate(ctx context.Context, date string) (bool, error)
}

type dateService struct {
}

func NewService() Service {
	return &dateService{}
}

func (d *dateService) Status(ctx context.Context) (string, error) {
	return "ok", nil
}

func (d *dateService) Get(ctx context.Context) (string, error) {
	now := time.Now()
	return now.Format(time.RFC3339), nil
}

func (d *dateService) Validate(ctx context.Context, date string) (bool, error) {
	_, err := time.Parse(time.RFC3339, date)
	if err != nil {
		return false, err
	}
	return true, nil
}
