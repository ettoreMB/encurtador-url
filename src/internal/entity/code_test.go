package entity_test

import (
	"encurtador_url/src/internal/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateCode(t *testing.T) {

	code, err := entity.NewCode("http://teste.code")

	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, code.Url, "http://teste.code")
	assert.Equal(t, len(code.Code), 6)

}

func Test_Should_not_createCode_without_url(t *testing.T) {
	_, err := entity.NewCode("")

	assert.NotNil(t, err)

}
