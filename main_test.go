package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test__simp_sum(t *testing.T) {
	res := simp_sum(2, 2)

	assert.Equal(t, 4, res)
}
