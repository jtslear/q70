package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Test struct {
	input         int
	desiredResult int
}

var td []Test

func TestTheQuestion(t *testing.T) {
	td = append(td, Test{
		1,
		19,
	})

	td = append(td, Test{
		2,
		28,
	})

	for _, v := range td {
		result := theQuestion(v.input)

		assert.Equalf(
			t,
			v.desiredResult,
			result,
			"Expected %v, got %v\n", v.desiredResult, result,
		)
	}
}
