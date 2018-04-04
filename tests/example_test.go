package tests

import (
	"testing"
	"github.com/stretchr/testify/assert"
)


func TestAdd(t *testing.T)  {
	expected := 5
	actual := Add(2, 3)
	assert.Equal(t, expected, actual)
}

func TestMul(t *testing.T)  {
	expected := 6
	actual := Mul(2, 3)
	assert.Equal(t, expected, actual)
}
