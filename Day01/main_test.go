package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_countIncreases(t *testing.T) {
	type args struct {
		input []int
	}
	tests := []struct {
		name           string
		args           args
		expectedResult int
	}{
		{
			name: "Empty input",
			args: args{
				input: make([]int, 0),
			},
			expectedResult: 0,
		},
		{
			name: "Simple input",
			args: args{
				input: []int{10, 20, 30},
			},
			expectedResult: 2,
		}, {
			name: "Challenge input",
			args: args{
				input: Challenge_input,
			},
			expectedResult: 1139,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := countIncreases(tt.args.input)
			assert.Equal(t, c, tt.expectedResult)
		})
	}
}
