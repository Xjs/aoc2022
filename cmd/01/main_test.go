package main

import (
	"io"
	"strings"
	"testing"
)

func Test_aoc1task1(t *testing.T) {
	t.Parallel()

	const exampleInput = `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`

	type args struct {
		r     io.Reader
		count int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{"example1", args{strings.NewReader(exampleInput), 1}, 24000, false},
		{"example2", args{strings.NewReader(exampleInput), 3}, 45000, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := aoc1(tt.args.r, tt.args.count)
			if (err != nil) != tt.wantErr {
				t.Errorf("aoc1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("aoc1() = %v, want %v", got, tt.want)
			}
		})
	}
}
