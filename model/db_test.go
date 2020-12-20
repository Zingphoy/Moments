package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO: 使用GoMock来mock异常情况

func TestPing(t *testing.T) {
	ret, err := Ping()
	assert.Equal(t, "success", ret)
	assert.Nil(t, err)
}
