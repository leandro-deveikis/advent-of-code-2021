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

func Test_buildWindows(t *testing.T) {
	type args struct {
		input []int
	}
	tests := []struct {
		name           string
		args           args
		expectedResult []int
	}{
		{
			name: "Empty input",
			args: args{
				input: make([]int, 0),
			},
			expectedResult: make([]int, 0),
		},
		{
			name: "Simple input",
			args: args{
				input: []int{10, 20, 30},
			},
			expectedResult: []int{60},
		},
		{
			name: "Simple input 2",
			args: args{
				input: []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263},
			},
			expectedResult: []int{607, 618, 618, 617, 647, 716, 769, 792},
		},
		/*	{
			name: "Challenge input",
			args: args{
				input: Challenge_input,
			},
			expectedResult: 1139,
		},  */
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := buildWindows(tt.args.input)
			assert.Equal(t, c, tt.expectedResult)
		})
	}
}
