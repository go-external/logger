package logger

import (
	"fmt"
	"testing"
)

type testCase struct {
	s string
	w Level
}

var tests = []testCase{
	{"trace", Trace},
	{"debug", Debug},
	{"info", Info},
	{"warn", Warn},
	{"error", Error},
	{"fatal", Fatal},
	{"wrong level", InvalidLevel},
	{"arn", InvalidLevel},
	{"eror", InvalidLevel},
	{"fetal", InvalidLevel},
}

func TestLevelParsing(t *testing.T) {
	for _, test := range tests {
		name := fmt.Sprintf("case(%s, %v)", test.s, test.w)
		t.Run(name, func(t *testing.T) {
			got, _ := parseLevel(test.s)
			if got != test.w {
				t.Errorf("got: %v, want: %v", got, test.w)
			}
		})
	}
}
