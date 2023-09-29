package pkg_test

import (
	"encurtador_url/src/pkg"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateRandomCodeWithSixChar(t *testing.T) {

	code := pkg.CodGenerator(6)

	assert.Equal(t, len(code), 6)

}
