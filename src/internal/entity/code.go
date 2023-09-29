package entity

import (
	"encurtador_url/src/pkg"
	"errors"
)

type Code struct {
	Url  string
	Code string
}

func NewCode(url string) (*Code, error) {

	c := Code{Url: url, Code: pkg.CodGenerator(6)}

	err := c.Validate()
	if err != nil {
		return nil, err
	}

	return &c, nil

}

func (c *Code) Validate() error {
	if c.Url == "" {
		return errors.New("url is empty")
	}

	if c.Code == "" {
		return errors.New("code is empty")
	}

	if len(c.Code) < 6 {
		return errors.New("code must have more then 6 letters")
	}

	return nil
}
