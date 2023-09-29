package repository

import (
	"encurtador_url/src/internal/entity"
)

type CodeRepositoryInMemory struct {
	Items []entity.Code
}

func NewCodeRepositoryInMemory() *CodeRepositoryInMemory {
	return &CodeRepositoryInMemory{}
}

func (r *CodeRepositoryInMemory) Save(code entity.Code) error {

	r.Items = append(r.Items, code)

	return nil
}

func (r *CodeRepositoryInMemory) Get(code string) (*entity.Code, error) {
	var result entity.Code
	for _, item := range r.Items {
		if item.Code == code {
			return &result, nil
		}
	}
	return &result, nil
}
