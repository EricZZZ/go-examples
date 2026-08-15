package products

import (
	"context"
	"database/sql"
	repo "ecom-api/internal/adapters/postgresql/sqlc"
	"errors"
	"fmt"
)

var (
	ErrProductNotFound = errors.New("product not found")
)

type Service interface {
	ListProducts(ctx context.Context) ([]repo.Product, error)
	FindProductByID(ctx context.Context, id int64) (repo.Product, error)
}

type svc struct {
	repo repo.Querier
}

func NewService(repo repo.Querier) Service {
	return &svc{repo: repo}
}

func (s *svc) ListProducts(ctx context.Context) ([]repo.Product, error) {
	return s.repo.ListProducts(ctx)
}

func (s *svc) FindProductByID(ctx context.Context, id int64) (repo.Product, error) {
	product, err := s.repo.FindProductByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return repo.Product{}, ErrProductNotFound
		}
		return repo.Product{}, fmt.Errorf("service errors")
	}
	return product, nil
}
