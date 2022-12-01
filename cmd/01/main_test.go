package main

import (
	"io"
	"strings"
	"testing"
)

func Test_aoc1task1(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{"example", args{strings.NewReader(
			`1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`,
		)}, 24000, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := aoc1task1(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("aoc1task1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("aoc1task1() = %v, want %v", got, tt.want)
			}
		})
	}
}
