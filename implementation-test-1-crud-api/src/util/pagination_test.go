package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLimit(t *testing.T) {
	for _, testCase := range []struct {
		PaginParam *PaginParam

		ExpectedLimit uint
	}{
		{
			PaginParam:    &PaginParam{Limit: 200},
			ExpectedLimit: 100,
		},
		{
			PaginParam:    &PaginParam{Limit: 40},
			ExpectedLimit: 40,
		},
	} {
		assert.Equal(t, testCase.ExpectedLimit, testCase.PaginParam.GetLimit(), "TestGetLimit")
	}
}

func TestGetOffset(t *testing.T) {
	for _, testCase := range []struct {
		PaginParam *PaginParam

		ExpectedOffset uint
	}{
		{
			PaginParam:     &PaginParam{Page: 0},
			ExpectedOffset: 0,
		},
		{
			PaginParam: &PaginParam{
				Page:  3,
				Limit: 100,
			},
			ExpectedOffset: 200,
		},
	} {
		assert.Equal(t, int(testCase.ExpectedOffset), int(testCase.PaginParam.GetOffset()), "TestGetOffset")
	}
}
