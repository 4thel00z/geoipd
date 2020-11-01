package libgeoip

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetGoEnv(t *testing.T) {

	env, err := GetGoEnv()
	assert.Nil(t, err)
	_, found := env["GOMOD"]
	assert.True(t, found)
}
