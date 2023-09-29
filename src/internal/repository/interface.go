package repository

import "encurtador_url/src/internal/entity"

type CodeRepositoryInterface interface {
	Save(code entity.Code) error
	Get(code string) (*entity.Code, error)
}
